var cont = 0;

function empezar() {
    for(i = 0; i < cont; i++)
        document.getElementById("pez" + i).style.filter = "grayscale(100%)";
    astilectron.send({name: "empezar", payload: {}})
    document.getElementById("boton").innerHTML = "Corriendo";
    document.getElementById("boton").disabled = true;
    astilectron.listen(function(message) {
        switch (message.name) {
        case "iteracion":
            document.getElementById("iteracion").innerHTML = message.payload.iteracion;
            document.getElementById("mejor-suma").innerHTML = message.payload.msuma;
            document.getElementById("mejor-fitness").innerHTML = message.payload.mfitness;
            document.getElementById("sub-tamanio").innerHTML = message.payload.subtamanio;
            document.getElementById("diferencia").innerHTML = message.payload.diferencia;
            for(i = 0; i < cont; i++)
                document.getElementById("pez" + i).style.filter = "grayscale(" + message.payload.peces[i] + "%)";
            for(i = 0; i < message.payload.subconjunto.length; i++) {
                if(message.payload.subconjunto[i])
                    document.getElementById("elemento" + i).style.background = "orange";
                else
                    document.getElementById("elemento" + i).style.background = "white";
            }
            break;
        case "terminado":
            document.getElementById("boton").innerHTML = "Volver a correr";
            document.getElementById("boton").disabled = false;
            break;
        }
    });
}

document.addEventListener('astilectron-ready', function() {
    astilectron.send({name: "ready", payload: {}})
    astilectron.listen(function(message) {
        switch (message.name) {
        case "ready":
            var peces = message.payload.npeces;
            var conjunto = message.payload.tconjunto;
            document.getElementById("conjunto").innerHTML = conjunto;
            document.getElementById("suma").innerHTML = message.payload.suma;
            document.getElementById("cardumen").innerHTML = peces;
            
            for(i = 0; i < Math.floor(Math.sqrt(peces)); i++) {
                var renglon = document.createElement("div");
                document.getElementById("peces").appendChild(renglon);
                for(j = 0; j < Math.floor(Math.sqrt(peces)); j++) {
                    var img = new Image(50,35);
                    img.id = "pez" + cont;                             
                    img.style.filter = "grayscale(100%)";
                    img.src = 'imagenes/pez.png';
                    renglon.appendChild(img);
                    cont++;
                }                
            }
            if(cont < peces - 1) {
                var renglon = document.createElement("div");
                document.getElementById("peces").appendChild(renglon);
            }
            
            while(cont < peces) {
                var img = new Image(50,35);
                img.id = "pez" + cont;                             
                img.style.filter = "grayscale(100%)";
                img.src = 'imagenes/pez.png';
                renglon.appendChild(img);
                cont++;
            }

            var textnode = document.createTextNode("1.0  ");
            document.getElementById("fitness").appendChild(textnode);
            for(i = 0; i < 10; i++) {
                var img = new Image(50,35);
                img.style.filter = "grayscale(" + (i * 10) + "%)";
                img.src = 'imagenes/pez.png';
                document.getElementById("fitness").appendChild(img);
            }
            textnode = document.createTextNode("  0.0");
            document.getElementById("fitness").appendChild(textnode);
            
            for(i = 0; i < conjunto; i++) {
                var elemento = document.createElement("div");
                elemento.id = "elemento" + i;                             
                elemento.className = "elemento";
                document.getElementById("subconjunto").appendChild(elemento);
            }
            break;
        }
    });
});
