$(document).ready(function(){

  $("#searchOptions").click(function(){
    $("#search-options").toggle();
  });

  $("#addRecipeOptions").click(function(){
    $("#add-recipe").toggle();
  })

  $("#search").click(function(){
    $("#clickTest").text("click");
    var dataPost="{"

    if (document.getElementById('search_name').value != null && document.getElementById('search_name').value != ""){
      dataPost=dataPost + '"Name" : "'+ document.getElementById('search_name').value+'",';
    }

    if (document.getElementById('search_type').value != "null" &&document.getElementById('search_type').value != ""){
      dataPost=dataPost + '"Type" : "'+ document.getElementById('search_type').value+'",';
    }

    if (document.getElementById('search_key').value != null && document.getElementById('search_key').value != ""){
      dataPost=dataPost + '"keywords" : "'+ document.getElementById('search_key').value+'",';
    }

    if (document.getElementById('search_rate').value != null && document.getElementById('search_rate').value != ""){
      dataPost=dataPost + '"rating" : "'+ document.getElementById('search_rate').value+'",';
    }

    $.post("/search/",
            dataPost,
            function(data, status) {
              $("#clickTest").text(status);
              console.log(data)
            }, "json"
          );
      });

  $("#addRecipe").click(function(){
    var dataPost='{';
    if (document.getElementById('add_name').value != null && document.getElementById('add_name').value != ""){
      dataPost=dataPost + '"Name" : "'+ document.getElementById('add_name').value+'",';
    }

    if (document.getElementById('add_type').value != "null" &&document.getElementById('add_type').value != ""){
      dataPost=dataPost + '"Type" : "'+ document.getElementById('add_type').value+'",';
    }

    if (document.getElementById('add_url').value != null && document.getElementById('add_url').value != ""){
      dataPost=dataPost + '"URL" : "'+ document.getElementById('add_url').value+'",';
    }

    if (document.getElementById('add_time').value != null && document.getElementById('add_time').value != ""){
      dataPost=dataPost + '"cooktime" : "'+ document.getElementById('add_time').value+'",';
    }

    if (document.getElementById('add_key').value != null && document.getElementById('add_key').value != ""){
      dataPost=dataPost + '"keywords" : "'+ document.getElementById('add_key').value+'",';
    }

    if (document.getElementById('add_rate').value != null && document.getElementById('add_rate').value != ""){
      dataPost=dataPost + '"rating" : "'+ document.getElementById('add_rate').value+'",';
    }

    console.log(dataPost);
    dataPost=dataPost.slice(0,-1);
    dataPost=dataPost+'}';
    $.post("/add/",
            dataPost,
            function(data, status) {
              $("#clickTest").text(status);
            }
          );
    console.log(dataPost);
  });

<!--$("#siteloader").html('<object data="http://allrecipes.com/recipe/234592/buffalo-chicken-stuffed-shells/"/>');-->
});
