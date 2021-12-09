import time
import tkinter
import tkinter.messagebox
from threading import Thread

def main():
    class DownloadTaskHandler(Thread):
        def run(self):
            time.sleep(10)
            tkinter.messagebox.showinfo('Info', "Download Done!")
            button1.config(state=tkinter.NORMAL)
        
    def download():
        button1.config(state=tkinter.DISABLED)
        DownloadTaskHandler(daemon=True).start()

    def show_about():
        tkinter.messagebox.showinfo("About", 'Author: QL')

    top = tkinter.Tk()
    top.title('Single Thread')
    top.geometry('200x150')
    top.wm_attributes('-topmost', 1)

    pannel = tkinter.Frame(top)
    button1 = tkinter.Button(pannel, text='Download', command=download)
    button1.pack(side='left')
    button2 = tkinter.Button(pannel, text='About', command=show_about)
    button2.pack(side='right')
    pannel.pack(side='bottom')
    tkinter.mainloop()

if __name__=='__main__':
    main()