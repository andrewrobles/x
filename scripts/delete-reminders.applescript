tell application "Reminders"
    repeat with reminderList in lists
        try
            delete reminderList
        end try
    end repeat
end tell