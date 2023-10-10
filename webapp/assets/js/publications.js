$('#nova-publication').on('submit', criarPublication);

$(document).on('click', '.curtir-publication', curtirPublication);
$(document).on('click', '.descurtir-publication', descurtirPublication);

$('#atualizar-publication').on('click', atualizarPublication);
$('.deletar-publication').on('click', deletarPublication);

function criarPublication(evento) {
    evento.preventDefault();

    $.ajax({
        url: "/publications",
        method: "POST",
        data: {
            title: $('#title').val(),
            conteudo: $('#conteudo').val(),
        }
    }).done(function() {
        window.location = "/home";
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao criar a publicação!", "error");
    })
}

function curtirPublication(evento) {
    evento.preventDefault();

    const elementoClicado = $(evento.target);
    const publicationId = elementoClicado.closest('div').data('publication-id');

    elementoClicado.prop('disabled', true);
    $.ajax({
        url: `/publications/${publicationId}/curtir`,
        method: "POST"
    }).done(function() {
        const contadorDeCurtidas = elementoClicado.next('span');
        const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text());

        contadorDeCurtidas.text(quantidadeDeCurtidas + 1);

        elementoClicado.addClass('descurtir-publication');
        elementoClicado.addClass('text-danger');
        elementoClicado.removeClass('curtir-publication');

    }).fail(function() {
        Swal.fire("Ops...", "Erro ao curtir a publicação!", "error");
    }).always(function() {
        elementoClicado.prop('disabled', false);
    });
}

function descurtirPublication(evento) {
    evento.preventDefault();

    const elementoClicado = $(evento.target);
    const publicationId = elementoClicado.closest('div').data('publication-id');

    elementoClicado.prop('disabled', true);
    $.ajax({
        url: `/publications/${publicationId}/descurtir`,
        method: "POST"
    }).done(function() {
        const contadorDeCurtidas = elementoClicado.next('span');
        const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text());

        contadorDeCurtidas.text(quantidadeDeCurtidas - 1);

        elementoClicado.removeClass('descurtir-publication');
        elementoClicado.removeClass('text-danger');
        elementoClicado.addClass('curtir-publication');

    }).fail(function() {
        Swal.fire("Ops...", "Erro ao descurtir a publicação!", "error");
    }).always(function() {
        elementoClicado.prop('disabled', false);
    });
}

function atualizarPublication() {
    $(this).prop('disabled', true);

    const publicationId = $(this).data('publication-id');
    
    $.ajax({
        url: `/publications/${publicationId}`,
        method: "PUT",
        data: {
            title: $('#title').val(),
            conteudo: $('#conteudo').val()
        }
    }).done(function() {
        Swal.fire('Sucesso!', 'Publicação criada com sucesso!', 'success')
            .then(function() {
                window.location = "/home";
            })
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao editar a publicação!", "error");
    }).always(function() {
        $('#atualizar-publication').prop('disabled', false);
    })
}

function deletarPublication(evento) {
    evento.preventDefault();

    Swal.fire({
        title: "Atenção!",
        text: "Tem certeza que deseja excluir essa publicação? Essa ação é irreversível!",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then(function(confirmacao) {
        if (!confirmacao.value) return;

        const elementoClicado = $(evento.target);
        const publication = elementoClicado.closest('div')
        const publicationId = publication.data('publication-id');
    
        elementoClicado.prop('disabled', true);
    
        $.ajax({
            url: `/publications/${publicationId}`,
            method: "DELETE"
        }).done(function() {
            publication.fadeOut("slow", function() {
                $(this).remove();
            });
        }).fail(function() {
            Swal.fire("Ops...", "Erro ao excluir a publicação!", "error");
        });
    })

}