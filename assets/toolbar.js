const editTextarea = document.getElementsByClassName('edit-form__textarea')[0]

function insertTextAtCursor(text, el = editTextarea) {
    const [start, end] = [el.selectionStart, el.selectionEnd];
    el.setRangeText(text, start, end, 'select');
}

function insertDate() {
    let date = new Date().toISOString().split('T')[0]
    insertTextAtCursor(date)
}
