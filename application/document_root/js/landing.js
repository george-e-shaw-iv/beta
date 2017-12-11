jQuery(document).ready(function(){
    var auth = 0;
    var landing = jQuery('.beta-landing');
    var content = jQuery('.beta-homepage-jumbotron');

    if(auth == 0) {
        var delay = 500;

        landing.find('span').each(function (i) {
          var _this = this;

          setTimeout(function () {
              jQuery(_this).animate({
                  opacity: 1
              }, 500);
          }, delay);

          delay += 500;
        });

        setTimeout(function () {
          landing.fadeOut("slow", function () {
              landing.attr("style", "display:none !important");
              content.fadeIn("slow");
          });
        }, 4000)
    } else {
        landing.attr("style", "display:none !important");
        content.attr("style", "display:block !important");
    }
});
