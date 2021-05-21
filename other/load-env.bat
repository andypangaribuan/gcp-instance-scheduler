@echo off
for /f "tokens=* eol=#" %%a in (.\%1) do (
    set %%a
)