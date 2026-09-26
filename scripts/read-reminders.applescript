tell application "Reminders"
	set markdown to ""
	repeat with reminderList in lists
		set reminderItems to reminders of reminderList
		if (count of reminderItems) > 0 then
			set markdown to markdown & "" & name of reminderList & linefeed
			repeat with r in reminderItems
				if completed of r then
					set checkbox to "- [x] "
				else
					set checkbox to "- [ ] "
				end if
				set markdown to markdown & checkbox & name of r & linefeed
			end repeat
			set markdown to markdown & linefeed
		end if
	end repeat
	return markdown
end tell
