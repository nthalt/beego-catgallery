document.addEventListener("DOMContentLoaded", () => {
  const tabButtons = document.querySelectorAll(".tab-button");
  const tabContents = document.querySelectorAll(".tab-content");
  const breedSelect = document.getElementById("breed-select");
  const voteButtons = document.querySelectorAll(".vote-button");
  const viewButtons = document.querySelectorAll(".view-button");
  const favCats = document.querySelector(".fav-cats");

  // Tab switching
  tabButtons.forEach((button) => {
    button.addEventListener("click", () => {
      tabButtons.forEach((btn) => btn.classList.remove("active"));
      tabContents.forEach((content) => content.classList.remove("active"));
      button.classList.add("active");
      document.getElementById(button.dataset.tab).classList.add("active");

      if (button.dataset.tab === "voting") {
        fetchRandomCat();
      } else if (button.dataset.tab === "breeds") {
        fetchBreeds();
      } else if (button.dataset.tab === "favs") {
        fetchFavoriteCats();
      }
    });
  });

  // Breed selection
  breedSelect.addEventListener("change", () => {
    const selectedBreed = breedSelect.value;
    if (selectedBreed) {
      fetchBreedInfo(selectedBreed);
    }
  });

  // Voting
  voteButtons.forEach((button) => {
    button.addEventListener("click", () => {
      const vote = button.dataset.vote;
      const catImage = document.querySelector(".cat-image img");
      if (catImage) {
        voteCat(catImage.dataset.id, vote);
      }
    });
  });

  // View switching
  viewButtons.forEach((button) => {
    button.addEventListener("click", () => {
      viewButtons.forEach((btn) => btn.classList.remove("active"));
      button.classList.add("active");
      favCats.className = "fav-cats " + button.dataset.view + "-view";
    });
  });

  // Initial load
  fetchRandomCat();
});

function fetchRandomCat() {
  fetch("/api/cats/random")
    .then((response) => response.json())
    .then((data) => {
      const catImage = document.querySelector(".cat-image");
      catImage.innerHTML = `<img src="${data.url}" alt="Random cat" data-id="${data.id}">`;
    })
    .catch((error) => console.error("Error fetching random cat:", error));
}

function fetchBreeds() {
  fetch("/api/breeds")
    .then((response) => response.json())
    .then((data) => {
      const breedSelect = document.getElementById("breed-select");
      breedSelect.innerHTML = '<option value="">Select a breed</option>';
      data.forEach((breed) => {
        const option = document.createElement("option");
        option.value = breed.id;
        option.textContent = breed.name;
        breedSelect.appendChild(option);
      });
    })
    .catch((error) => console.error("Error fetching breeds:", error));
}

function fetchBreedInfo(breedId) {
  fetch(`/api/breeds/${breedId}`)
    .then((response) => response.json())
    .then((data) => {
      const breedImageSlider = document.querySelector(".breed-image-slider");
      breedImageSlider.innerHTML = "";

      data.forEach((image) => {
        const img = document.createElement("img");
        img.src = image.url;
        img.alt = "Breed image";
        breedImageSlider.appendChild(img);
      });

      // Fetch breed details
      fetch(`/api/breeds`)
        .then((response) => response.json())
        .then((breeds) => {
          const breed = breeds.find((b) => b.id === breedId);
          if (breed) {
            const breedDescription =
              document.querySelector(".breed-description");
            breedDescription.innerHTML = `
              <h3>${breed.name}</h3>
              <p>${breed.description}</p>
              <p><strong>Temperament:</strong> ${breed.temperament}</p>
              <p><strong>Origin:</strong> ${breed.origin}</p>
              <p><strong>Life Span:</strong> ${breed.life_span} years</p>
              <p><strong>Weight:</strong> ${breed.weight.metric} kg</p>
            `;
          }
        })
        .catch((error) =>
          console.error("Error fetching breed details:", error)
        );
    })
    .catch((error) => console.error("Error fetching breed info:", error));
}

// new
function voteCat(imageId, value) {
  if (value === "1") {
    addFavorite(imageId);
  } else {
    fetch("/api/votes", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ image_id: imageId, value: parseInt(value) }),
    })
      .then((response) => response.json())
      .then((data) => {
        console.log("Vote submitted:", data);
        fetchRandomCat(); // Load a new cat after voting
      })
      .catch((error) => console.error("Error voting:", error));
  }
}

function addFavorite(imageId) {
  fetch("/api/favourites", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ image_id: imageId }),
  })
    .then((response) => response.json())
    .then((data) => {
      console.log("Favorite added:", data);
      alert("Cat added to favorites!");
      fetchRandomCat(); // Load a new cat after favoriting
    })
    .catch((error) => console.error("Error adding favorite:", error));
}

function fetchFavoriteCats() {
  fetch("/api/favourites")
    .then((response) => response.json())
    .then((data) => {
      const favCats = document.querySelector(".fav-cats");
      favCats.innerHTML = "";
      data.forEach((cat) => {
        const catElement = document.createElement("div");
        catElement.className = "fav-cat";
        catElement.innerHTML = `
          <img src="${cat.image.url}" alt="Favorite cat">
          <p>ID: ${cat.id}</p>
        `;
        favCats.appendChild(catElement);
      });
    })
    .catch((error) => console.error("Error fetching favorite cats:", error));
}
