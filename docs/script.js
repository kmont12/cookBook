$(document).ready(function(){
  var searchResults, website;
  $("#searchOptions").click(function(){
    $("#search-options").toggle();
    $("#add-recipe").hide();
    $("#list-results").empty();
    $('#siteloader').hide();
  });

  $("#addRecipeOptions").click(function(){
    $("#add-recipe").toggle();
    $("#search-options").hide();
    $("#list-results").empty();
    $('#siteloader').hide();
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
      dataPost=dataPost + '"Keywords" : "'+ document.getElementById('search_key').value+'",';
    }

    if (document.getElementById('search_rate').value != null && document.getElementById('search_rate').value != ""){
      dataPost=dataPost + '"Rating" : "'+ document.getElementById('search_rate').value+'",';
    }

    dataPost=dataPost.slice(0,-1);
    dataPost=dataPost+'}';
    console.log(dataPost);
    $.post("/search/",
            dataPost,
            function(data, status) {
              console.log(data);
              makeList(data);
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
      dataPost=dataPost + '"Cooktime" : '+ document.getElementById('add_time').value+',';
    }

    if (document.getElementById('add_key').value != null && document.getElementById('add_key').value != ""){
      dataPost=dataPost + '"Keywords" : "'+ document.getElementById('add_key').value+'",';
    }

    if (document.getElementById('add_rate').value != null && document.getElementById('add_rate').value != ""){
      dataPost=dataPost + '"Rating" : '+ document.getElementById('add_rate').value+',';
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

  function makeList(jsonData){
    $("#search-options").hide();
    $("#list-results").empty();
    $("#list-results").show();
    searchResults=jsonData;
    $.each(jsonData, function(key, value) {
      $('#list-results').append('<li id="' + key + '" onClick="listClickListener(\''+key+'\')">' + key + '</li>');
      document.getElementById(key).addEventListener("click", function(){listClickListener(key);}, false);
    });
  }


  function listClickListener(id){
   console.log(id);
   $.each(searchResults, function(key, value){
     if (id == key){
       website=value.url;
       //console.log(website);
     }
     $("#siteloader").html('<object data="'+website+'"/>');
     $('#siteloader').show();
     $('#list-results').hide();
   });
   return false;
  }

});
