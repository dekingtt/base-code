import json
import ipaddress


def s_rule_detection_APT_process_chain(line_j):
    alert = {}
    alert["name"] = "APT process chain detected"
    alert["timestamp"] = line_j["unixTime"]
    alert["agent"] = line_j["hostIdentifier"]
    alert["events"] = [line_j]
    alert["description"] = "Office start suspicious child process, then child process make internet connection"
    with open("alerts.json", "a") as f:
        f.write(json.dumps(alert)+"\n")
    print("s_rule_detection_APT_process_chain")


def c_rule_possible_port_scan(lines):
    conn = {}
    event = {}
    for line in lines:
        line_j = json.loads(line)
        if line_j["name"] == "pack/Global/windows_process_outbound_connections":
            if ipaddress.ip_address(line_j["columns"]["destination_address"]).is_private:
                if line_j["hostIdentifier"] in conn:
                    if line_j["columns"]["destination_address"] in conn[line_j["hostIdentifier"]]:
                        conn[line_j["hostIdentifier"]][line_j["columns"]["destination_address"]].append(line_j["columns"]["destination_port"])
                    else:

                        event[line_j["columns"]["destination_address"]] = line_j
                        conn[line_j["hostIdentifier"]][line_j["columns"]["destination_address"]] = []
                        conn[line_j["hostIdentifier"]][line_j["columns"]["destination_address"]].append(line_j["columns"]["destination_port"])
                else:
                    conn[line_j["hostIdentifier"]] = {}
                    event[line_j["columns"]["destination_address"]] = line_j
                    conn[line_j["hostIdentifier"]][line_j["columns"]["destination_address"]] = []
                    conn[line_j["hostIdentifier"]][line_j["columns"]["destination_address"]].append(line_j["columns"]["destination_port"])
    for host in conn:
        for ip in conn[host]:
            if len(set(conn[host][ip])) > 100:
                alert = {}
                alert["name"] = "possible port scan detected"
                alert["timestamp"] = event[ip]["unixTime"]
                alert["agent"] = host
                alert["events"] = [event[ip]]
                alert["description"] = "{} is scanning ports of internal server {}".format(event[ip]["columns"]["source_address"], ip)
                with open("alerts.json", "a") as f:
                    f.write(json.dumps(alert)+"\n")
                print("c_rule_possible_port_scan")
                break


simple_rules = {"pack/Global/detection_APT_process_chain": s_rule_detection_APT_process_chain
        }


complex_rules = {"possible_port_scan": c_rule_possible_port_scan
        }


def simple_rule_match():
    with open("../logs/osqueryd.results.log", "r") as f:
        while True:
            line = f.readline()
            if not line:
                break
            line_j = json.loads(line)
            if line_j["name"] in simple_rules and line_j["action"] == "added":
                simple_rules[line_j["name"]](line_j)


def complex_rule_match():
    with open("../logs/osqueryd.results.log", "r") as f:
        lines = f.readlines()
    for item in complex_rules:
        complex_rules[item](lines)


def main():
    simple_rule_match()
    complex_rule_match()

if __name__ == "__main__":
    main()
