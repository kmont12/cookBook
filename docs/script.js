$(document).ready(function(){

  $("#searchOptions").click(function(){
    $("#search-options").toggle();
  });

  $("#addRecipeOptions").click(function(){
    $("#add-recipe").toggle();
  })

  $("#search").click(function(){
    $("#clickTest").text("click");
    var path="http://localhost:8888/search";

    var response;
    $.getJSON( path, function(data) {
          console.log( "success" );
          response = data;
          console.log(response);
        });
      });

  $("#addRecipe").click(function(){
    console.log("added");
    $.post("http://localhost:8888/add/",
            '{"ID" : 2, "Name" : "Dish"}',
            function(data, status) {
              $("#clickTest").text(status);
            }
          );
    console.log("after");
  });

<!--$("#siteloader").html('<object data="http://allrecipes.com/recipe/234592/buffalo-chicken-stuffed-shells/"/>');-->
});
