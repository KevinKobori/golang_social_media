$('#signIn').on('submit', fazerSignIn);

function fazerSignIn(evento) {
    evento.preventDefault();

    $.ajax({
        url: "/signIn",
        method: "POST",
        data: {
            email: $('#email').val(),
            senha: $('#senha').val(),
        }
    }).done(function() {
        window.location = "/home";
    }).fail(function() {
        Swal.fire("Ops...", "Usuário ou senha incorretos!", "error");
    });
}