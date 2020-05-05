// This adds 'condensed' toggles class on <body> depending on visibility of
// element with 'navbar-trigger' class. It's used to toggle 'mini-navbar'.
var navbar_trigger = document.querySelector('.navbar-trigger');
if (navbar_trigger) {
  (new IntersectionObserver(function(e,o){
    if (e[0].intersectionRatio > 0){
      $("body").removeClass("condensed");
    } else {
      $("body").addClass("condensed");
    };
  })).observe(navbar_trigger);
}
