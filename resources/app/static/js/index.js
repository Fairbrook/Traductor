
function onProcess(){
    const content = document.querySelector('#res')
    const textArea = document.querySelector('#content')
    const message = {}
    message.name="process"
    message.payload=textArea.value
    astilectron.sendMessage(message, function(message){
        // console.log(JSON.stringify(message.paylad));
        if(!Array.isArray( message.payload)){
            content.innerHTML = message.payload
        }
        const table = document.createElement("table")
        // table.style.borderWidth="2px"
        // table.style.borderColor="black"
        // table.style.borderStyle="solid"
        const tableBody = document.createElement("tbody")
        const headerRow = document.createElement("thead")
        let header = document.createElement('th')
        header.innerHTML = "Lexoma"
        headerRow.appendChild(header)
        header = document.createElement('th')
        header.innerHTML = "Tipo"
        headerRow.appendChild(header)
        table.appendChild(headerRow)
        table.appendChild(tableBody)
        message.payload.forEach((segment)=>{
            const row = document.createElement('tr')
            let cell = document.createElement('td')
            cell.innerHTML = segment.lexema
            row.appendChild(cell)
            cell = document.createElement('td')
            cell.innerHTML = segment.state_name
            row.appendChild(cell)
            tableBody.appendChild(row)
        })
        content.replaceChild(table, content.childNodes[0])
    })
}