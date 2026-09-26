tell application "Reminders"
    repeat with reminderList in lists
        delete every reminder of reminderList
    end repeat
end tell