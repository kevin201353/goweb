function mSign_signMark_app() {
    var wrapper = document.getElementById("mSign_signMark_signature_pad"),
        clearButton = document.getElementById("mSign_signMark_clear_out"),
        canvas = wrapper.querySelector("canvas"),
        sure = document.getElementsByClassName("mSign_signMark_footer_sure"),
        signaturePad;

    function resizeCanvas() {

        var ratio =  Math.max(window.devicePixelRatio || 1, 1);
        canvas.width = canvas.offsetWidth * ratio;
        canvas.height = canvas.offsetHeight * ratio;
        canvas.getContext("2d").scale(ratio, ratio);
    }

    window.onresize = resizeCanvas;
    resizeCanvas();

    signaturePad = new SignaturePad(canvas,{
        minWidth: 1,
        maxWidth: 2,
        penColor: "rgb(0,0,0)"
    });

    clearButton.addEventListener("click", function (event) {
        signaturePad.clear();
    });
}

