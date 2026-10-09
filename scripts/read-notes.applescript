tell application "Reminders"
    set allNotes to ""
    repeat with r in reminders of list "Reminders"
        set reminderNotes to body of r
        if reminderNotes is not missing value and reminderNotes is not "" then
            if allNotes is not "" then
                set allNotes to allNotes & linefeed
            end if
            set allNotes to allNotes & reminderNotes
        end if
    end repeat
    return allNotes
end tell