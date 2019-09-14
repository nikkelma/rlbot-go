rem Remove old generated files
Remove-Item .\flat\ -Recurse -Force

rem Generate go files from flatbuffer schema
.\flatbuffer\bin\flatc.exe --go .\flatbuffer\schema\rlbot.fbs

rem Clean up intermediate folder
Move-Item .\rlbot\flat\ .\flat\

rem Remove old generated files
Remove-Item .\rlbot -Recurse

rem Apply go fmt
go fmt .\flat\...
