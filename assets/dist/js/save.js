$(document).ready(function() {
    $('#save').on('click', function() {
        var notes = document.getElementById('note');
        var fd = new FormData(notes);
        req = $.ajax({
            url: '/save',
            type: 'POST',
            data: { src: src, cost: cost, name: name }
        });
        req.done(function(data) {
            $('#lilcart').text(data.cart_ammount);
        });
    });
});
