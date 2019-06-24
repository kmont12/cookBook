var currentRecipeID;
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
    if (document.getElementById('search_cook').value != null && document.getElementById('search_cook').value != ""){
      dataPost=dataPost + '"Cooktime" : "'+ document.getElementById('search_cook').value+'",';
    }

    dataPost=dataPost.slice(0,-1);
    dataPost=dataPost+'}';
//    console.log(dataPost);
    $.get("/recipe/search/",
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
    $.post("/recipe/add/",
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
      $('#list-results').append('<li id="' + key + '" onClick="listClickListener(this.id)">' + key + '</li>');
      document.getElementById(key).addEventListener("click", function(){listClickListener(key);}, false);
    });
  }


  function listClickListener(id){
   console.log(id);
   $.each(searchResults, function(key, value){
     if (id == key){
       website=value.url;
       currentRecipeID = value.id
       //console.log(website);
     }
     $("#siteloader").html('<object data="'+website+'"/>');
     $('#siteloader').show();
     $('#list-results').hide();
   });
   return false;
  }

});

function openNoteSidebar() {
  document.getElementById("noteSidebar").style.width = "30%";
  document.getElementById("main").style.marginRight = "30%";
}

function makeNotesList(data) {
  $.each(jsonData, function(key, value) {
    $('#notesList').append('<li id="' + key + '" dblclick="listClickListener(this.id)">' + key + '</li>');
    document.getElementById(key).addEventListener("click", function(){listClickListener(key);}, false);
  });
}

function closeNoteSidebar() {
  document.getElementById("noteSidebar").style.width = "0";
  document.getElementById("main").style.marginRight = "0";
}

function prepareForDeletion(id) {
  console.log("Deleting Note " + id )
  document.getElementById(id).style.display = 'none';
}