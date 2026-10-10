use framework "Foundation"
use scripting additions

property serializer : missing value

on run argv
    if (count of argv) is not 1 then
        error "Usage: osascript read.test.applescript /absolute/path/to/read.scpt"
    end if

    set scriptPath to item 1 of argv
    set scriptFile to POSIX file scriptPath
    set serializer to load script scriptFile
    return my testRemindersToJSON()
end run


on parseJSON(jsonText)
    set textValue to current application's NSString's stringWithString:jsonText
    set jsonData to textValue's dataUsingEncoding:(current application's NSUTF8StringEncoding)

    set {jsonValue, jsonError} to current application's NSJSONSerialization's JSONObjectWithData:jsonData options:0 |error|:(reference)

    if jsonValue is missing value then
        error ("Parse JSON: " & (jsonError's localizedDescription() as text))
    end if

    return jsonValue
end parseJSON


on assertJSON(testName, inputData, expectedJSON)
    set actualJSON to serializer's reminderListsToJSON(inputData)
    set actualValue to my parseJSON(actualJSON)
    set expectedValue to my parseJSON(expectedJSON)

    if not (actualValue's isEqualToArray:expectedValue) then
        error ("FAIL: " & testName & linefeed & "Actual JSON: " & actualJSON)
    end if
end assertJSON


on testRemindersToJSON()
    my assertJSON("No lists", {}, "[]")

    my assertJSON("Empty lists omitted", ¬
        {{listName:"Empty", reminderItems:{}}}, ¬
        "[]")

    set sampleData to {¬
        {listName:"Empty", reminderItems:{}}, ¬
        {listName:"Groceries", reminderItems:{¬
            {reminderName:"Buy milk", isCompleted:false}, ¬
            {reminderName:"Buy coffee", isCompleted:true}}}, ¬
        {listName:"Work", reminderItems:{¬
            {reminderName:"Review PR", isCompleted:false}}}}

    set expectedJSON to "[{\"name\":\"Groceries\",\"reminders\":[{\"name\":\"Buy milk\",\"completed\":false},{\"name\":\"Buy coffee\",\"completed\":true}]},{\"name\":\"Work\",\"reminders\":[{\"name\":\"Review PR\",\"completed\":false}]}]"

    my assertJSON("Completion flags and list order", sampleData, expectedJSON)

    set escapedTitle to "Quote \" and slash \\ and newline" & linefeed & "next"
    set sampleData to {{listName:"List \"A\"", reminderItems:{¬
        {reminderName:escapedTitle, isCompleted:false}}}}

    set expectedJSON to "[{\"name\":\"List \\\"A\\\"\",\"reminders\":[{\"name\":\"Quote \\\" and slash \\\\ and newline\\nnext\",\"completed\":false}]}]"

    my assertJSON("JSON escaping", sampleData, expectedJSON)

    return "PASS: 4 tests"
end testRemindersToJSON