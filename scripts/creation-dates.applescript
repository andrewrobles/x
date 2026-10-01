on getCreationDates(listName)
	tell application "Reminders"
		set output to ""
		
		repeat with r in reminders of list listName
			set createdAt to creation date of r
			
			if output is not "" then
				set output to output & ", "
			end if
			
			set output to output & (createdAt as string)
		end repeat
		
		return output
	end tell
end getCreationDates


on run argv
	set listName to item 1 of argv
	return getCreationDates(listName)
end run