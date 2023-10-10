$('#parar-de-follow').on('click', pararDeFollow);
$('#follow').on('click', follow);
$('#edit-user').on('submit', editar);
$('#atualizar-senha').on('submit', atualizarSenha);
$('#deletar-user').on('click', deletarUser);

function pararDeFollow() {
    const userId = $(this).data('user-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `/users/${userId}/parar-de-follow`,
        method: "POST"
    }).done(function() {
        window.location = `/users/${userId}`;
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao parar de follow o usuário!", "error");
        $('#parar-de-follow').prop('disabled', false);
    });
}

function follow() {
    const userId = $(this).data('user-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `/users/${userId}/follow`,
        method: "POST"
    }).done(function() {
        window.location = `/users/${userId}`;
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao follow o usuário!", "error");
        $('#follow').prop('disabled', false);
    });
}

function editar(evento) {
    evento.preventDefault();

    $.ajax({
        url: "/edit-user",
        method: "PUT",
        data: {
            nome: $('#nome').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
        }
    }).done(function() {
        Swal.fire("Sucesso!", "Usuário atualizado com sucesso!", "success")
            .then(function() {
                window.location = "/perfil";
            });
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao atualizar o usuário!", "error");
    });
}

function atualizarSenha(evento) {
    evento.preventDefault();

    if ($('#nova-senha').val() != $('#confirmar-senha').val()) {
        Swal.fire("Ops...", "As senhas não coincidem!", "warning");
        return;
    }

    $.ajax({
        url: "/atualizar-senha",
        method: "POST",
        data: {
            atual: $('#senha-atual').val(),
            nova: $('#nova-senha').val()
        }
    }).done(function() {
        Swal.fire("Sucesso!", "A senha foi atualizada com sucesso!", "success")
            .then(function() {
                window.location = "/perfil";
            })
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao atualizar a senha!", "error");
    });
}

function deletarUser() {
    Swal.fire({
        title: "Atenção!",
        text: "Tem certeza que deseja apagar a sua conta? Essa é uma ação irreversível!",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then(function(confirmacao) {
        if (confirmacao.value) {
            $.ajax({
                url: "/deletar-user",
                method: "DELETE"
            }).done(function() {
                Swal.fire("Sucesso!", "Seu usuário foi excluído com sucesso!", "success")
                    .then(function() {
                        window.location = "/logout";
                    })
            }).fail(function() {
                Swal.fire("Ops...", "Ocorreu um erro ao excluir o seu usuário!", "error");
            });
        }
    })
}