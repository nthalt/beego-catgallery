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
        // voteCat(catImage.dataset.id, vote);
        addFavorite(catImage.dataset.id);
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

// function for fetching all cat breeds and showing them as options in selector
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

// function for fetching all cats of the selected breed
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
        // img.classList.add("carousel-slide");
        breedImageSlider.appendChild(img);
      });

      // Initialize carousel after images are loaded
      initCarousel();

      // Fetch breed details
      fetch(`/api/breeds`)
        .then((response) => response.json())
        .then((breeds) => {
          const breed = breeds.find((b) => b.id === breedId);
          if (breed) {
            const breedDescription =
              document.querySelector(".breed-description");
            breedDescription.innerHTML = `
              <h3>
                <span>${breed.name}</span>
                <span class="breed-description-text"><i><bold>(${breed.origin})</bold></i></span>
                <span class="breed-description-text"><i>${breed.id}</i></span>
              </h3>
              <p class="breed-description-text">${breed.description}</p>
              <div class="wiki-link"><a href="${breed.wikipedia_url}">WIKIPEDIA</a></div>
            `;
          }
          // <span class="text-sm text-gray-500">(${breed.origin})</span>
          // <span class="text-sm italic font-light text-gray-500">${breed.id}</span>

          // <p><strong>Temperament:</strong> ${breed.temperament}</p>
          // <p><strong>Origin:</strong> ${breed.origin}</p>
          // <p><strong>Life Span:</strong> ${breed.life_span} years</p>
          // <p><strong>Weight:</strong> ${breed.weight.metric} kg</p>
        })
        .catch((error) =>
          console.error("Error fetching breed details:", error)
        );
    })
    .catch((error) => console.error("Error fetching breed info:", error));
}

function initCarousel() {
  const carousel = document.querySelector(".breed-image-slider");
  const images = carousel.querySelectorAll("img");
  let currentIndex = 0;

  images.forEach((img, index) => {
    img.style.display = index === 0 ? "block" : "none"; // Show the first image
  });

  setInterval(() => {
    images[currentIndex].style.display = "none"; // Hide the current image
    currentIndex = (currentIndex + 1) % images.length; // Move to the next image
    images[currentIndex].style.display = "block"; // Show the next image
  }, 2000); // Change image every 3 seconds
}

// new
function addFavorite(imageId) {
  const data = {
    image_id: imageId,
    // We no longer need to send sub_id from the client
  };

  console.log("Sending favorite data:", data);

  fetch("/api/favourites", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  })
    .then((response) => response.json())
    .then((data) => {
      console.log("Response from server:", data);
      if (data.error) {
        console.error("Error adding favorite:", data.error);
        alert("Failed to add cat to favorites: " + data.error);
      } else {
        console.log("Favorite added:", data);
        // alert("Cat added to favorites!");
        fetchRandomCat();
      }
    })
    .catch((error) => {
      console.error("Error adding favorite:", error);
      alert("Failed to add cat to favorites: " + error.message);
    });
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
        `;
        // <p>ID: ${cat.id}</p>
        favCats.appendChild(catElement);
      });
    })
    .catch((error) => console.error("Error fetching favorite cats:", error));
}

// function voteCat(imageId, value) {
//   if (value === "1") {
//     addFavorite(imageId);
//   } else {
//     fetch("/api/votes", {
//       method: "POST",
//       headers: {
//         "Content-Type": "application/json",
//       },
//       body: JSON.stringify({ image_id: imageId, value: parseInt(value) }),
//     })
//       .then((response) => response.json())
//       .then((data) => {
//         console.log("Vote submitted:", data);
//         fetchRandomCat(); // Load a new cat after voting
//       })
//       .catch((error) => console.error("Error voting:", error));
//   }
// }
