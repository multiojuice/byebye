# byebye
An configurable command line tool written in Go to end or kill any set of processes/apps running on the computer
This tool will end any list of process you put in the .byebyerc file with the specified type of end: interrupt, hangup, terminate, kill. 

just put this binary in you $PATH, so maybe in /usr/bin or something of that sort. 

then setup you .byebyerc in your homedir
use `byebye help` to see the best way to do that. 

Disclaimer : Be careful and only use this if you know your config will not kill something that will mess your computer up.

Example ~/.byebyerc
interrupt chrome
interrupt telegram
interrupt code
termintate tmux
