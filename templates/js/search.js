const searchInput = document.getElementById("searchInput");
const suggestionsContainer = document.getElementById("suggestions");
const artistList = document.getElementById("groupsList");
var flag = false
const list000 = document.getElementById("searchContainer");
const copyArtistList = artistList.cloneNode(true).innerHTML;
const images = list000.getElementsByTagName("img");
const image2 = [...images];
searchInput.addEventListener("input", () => {
    const searchText = searchInput.value;
    fetchSuggestions(searchText);
});
function fetchSuggestions(searchText) {
    if (searchText.trim() === "") {
        suggestionsContainer.innerHTML = "";
        artistList.innerHTML = copyArtistList;
        return;
    }
    // Send request to the backend to get suggestions
    fetch(`/search?searchText=${searchText}`)
        .then(response => response.json())
        .then(data => displaySuggestions(data))
        .catch(error => console.error(error));
}
function displaySuggestions(suggestions) {
   if (flag) {
    flag = false
    suggestionsContainer.innerHTML = "";
   } else  {
    console.log(suggestions)
    if (suggestions["suggestions"].length < 1) {
    console.log("len is zero");
    suggestionsContainer.innerHTML = ""
    const notFoundElement = document.createElement("option");
    notFoundElement.textContent = "not found";
    artistList.innerHTML = "";
    artistList.appendChild(notFoundElement);
    return
  }
    suggestionsContainer.innerHTML = "";
    for (let suggestion of suggestions["suggestions"]) {
        const suggestionDiv = document.createElement("option");
        suggestionDiv.classList.add("suggestion");
        suggestionDiv.textContent = suggestion;
        suggestionDiv.addEventListener("click", () => {
            const startIndex = suggestion.indexOf("-");
            const trimmedStr = startIndex !== -1 ? suggestion.substring(0, startIndex) : suggestion;
            searchInput.value = trimmedStr;
            suggestionsContainer.innerHTML = "";
            flag = true;
            fetchSuggestions(trimmedStr)
        });
        suggestionsContainer.appendChild(suggestionDiv);    
    }
   }
    
    artistList.innerHTML = "";
    for (let key in suggestions["artists"]) {
        const a = document.createElement("a");
        const img = document.createElement("img");
        const h2 = document.createElement("h2");
        h2.style.textAlign = "center";
        h2.style.fontSize = "20px"
        h2.style.color = "white"
        img.style.paddingRight = "20px"
        a.href = "/groups/" + key;
        a.style.textDecoration = "none"
        img.src = image2[key-1].getAttribute("src");
        h2.textContent = suggestions["artists"][key];
        a.appendChild(img);
        a.appendChild(h2);
        let container = document.createElement("div")
        container.style.marginBottom = "20px"
        container.appendChild(a)
        artistList.appendChild(container);
     }
}