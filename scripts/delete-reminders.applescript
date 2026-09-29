tell application "Reminders"
    set listNames to name of every list
    repeat with listName in listNames
        try
            delete list listName
        end try
    end repeat
end tell