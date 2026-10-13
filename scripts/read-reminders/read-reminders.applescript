use framework "Foundation"

on run
    return my remindersToJSON()
end run


on remindersToJSON()
    return my reminderListsToJSON(my readReminderLists())
end remindersToJSON


on readReminderLists()
    set collectedLists to {}

    tell application "Reminders"
        repeat with reminderList in lists
            set collectedReminders to {}

            repeat with r in reminders of reminderList
                set reminderTitle to name of r
                set reminderCompleted to completed of r
                set reminderBody to body of r

                set end of collectedReminders to {¬
                    reminderName:reminderTitle, ¬
                    reminderNotes:reminderBody, ¬
                    isCompleted:reminderCompleted}
            end repeat

            set reminderListName to name of reminderList

            set end of collectedLists to {¬
                listName:reminderListName, ¬
                reminderItems:collectedReminders}
        end repeat
    end tell

    return collectedLists
end readReminderLists


on reminderListsToJSON(reminderLists)
    set listData to current application's NSMutableArray's array()

    repeat with listRecord in reminderLists
        set reminderRecords to reminderItems of listRecord

        if (count of reminderRecords) > 0 then
            set reminderData to current application's NSMutableArray's array()

            repeat with reminderRecord in reminderRecords
                set title to reminderName of reminderRecord
                set notesValue to reminderNotes of reminderRecord
                set doneValue to current application's NSNumber's numberWithBool:(isCompleted of reminderRecord)

                set entry to current application's NSDictionary's dictionaryWithObjects:{title, notesValue, doneValue} forKeys:{"name", "notes", "completed"}
                reminderData's addObject:entry
            end repeat

            set entry to current application's NSDictionary's dictionaryWithObjects:{listName of listRecord, reminderData} forKeys:{"name", "reminders"}
            listData's addObject:entry
        end if
    end repeat

    set {jsonData, jsonError} to current application's NSJSONSerialization's dataWithJSONObject:listData options:0 |error|:(reference)

    if jsonData is missing value then
        error ("Serialize reminders: " & (jsonError's localizedDescription() as text))
    end if

    set jsonText to current application's NSString's alloc()'s initWithData:jsonData encoding:(current application's NSUTF8StringEncoding)
    return jsonText as text
end reminderListsToJSON