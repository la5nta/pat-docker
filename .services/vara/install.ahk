#NoTrayIcon
#SingleInstance
SetTitleMatchMode, 2
Run, %1% /SILENT, C:\
; Wait for "installed successfully" window and click the OK button
WinWait, VARA Setup
ControlClick, Button1, VARA Setup
WinWaitClose
