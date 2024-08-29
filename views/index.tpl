<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>The Cat API - Cats as a service</title>
    <link rel="stylesheet" type="text/css" href="/static/css/style.css" />
    <link
      rel="stylesheet"
      href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.3/css/all.min.css"
    />
  </head>
  <body>
    <header>
      <nav>
        <ul>
          <li><a href="#pricing">PRICING</a></li>
          <li><a href="#documentation">DOCUMENTATION</a></li>
          <li><a href="#more-apis">MORE APIS</a></li>
          <li><a href="#showcase">SHOWCASE</a></li>
        </ul>
      </nav>
    </header>
    <main>
      <section class="hero">
        <h1>The Cat API</h1>
        <h2>Cats as a service.</h2>
        <p>Because everyday is a Caturday.</p>
        <p>An API all about cat.</p>
        <p>60k+ Images. Breeds. Facts.</p>
        <button
          id="getApiKey"
          onclick="window.open('https://thecatapi.com/signup')"
        >
          GET YOUR API KEY
        </button>
        <button
          id="readGuides"
          onclick="window.open('https://developers.thecatapi.com/view-account/ylX4blBYT9FaoVd6OhvR?report=FJkYOq9tW')"
        >
          READ OUR GUIDES
        </button>
      </section>
      <section class="cat-card">
        <div class="tab-buttons">
          <button class="tab-button active" data-tab="voting">
            <i class="fas fa-poll-h"></i> Voting
          </button>
          <button class="tab-button" data-tab="breeds">
            <i class="fas fa-search"></i> Breeds
          </button>
          <button class="tab-button" data-tab="favs">
            <i class="fas fa-heart"></i> Favs
          </button>
        </div>
        <div id="voting" class="tab-content active">
          <div class="cat-image"></div>
          <div class="voting-buttons">
            <button class="vote-button" data-vote="1">
              <i class="fas fa-heart"></i>
            </button>
            <button class="vote-button" data-vote="0">
              <i class="fas fa-thumbs-up"></i>
            </button>
            <button class="vote-button" data-vote="0">
              <i class="fas fa-thumbs-down"></i>
            </button>
          </div>
        </div>
        <div id="breeds" class="tab-content">
          <select id="breed-select">
            <option value="">Select a breed</option>
          </select>
          <div class="breed-info">
            <div class="breed-image-slider"></div>
            <div class="breed-description"></div>
          </div>
        </div>
        <div id="favs" class="tab-content">
          <div class="view-buttons">
            <button class="view-button active" data-view="grid">
              <i class="fas fa-th"></i>
            </button>
            <button class="view-button" data-view="list">
              <i class="fas fa-list"></i>
            </button>
          </div>
          <div class="fav-cats"></div>
        </div>
      </section>
      <br /><br /><br />
    </main>
    <footer>
      <div class="footer-container">
        <!-- <div class="footer-section logo">
          <img src="logo.png" alt="Cat API" />
        </div> -->
        <div class="footer-section">
          <h4>Details</h4>
          <div class="footer-links">
            <a href="#">About Us</a>
          </div>
        </div>
        <div class="footer-section">
          <h4>Community and Support</h4>
          <div class="footer-links">
            <a href="#">Discord</a>
            <a href="#">Forum</a>
          </div>
        </div>
        <div class="footer-section">
          <h4>Our other APIs</h4>
          <div class="footer-links">
            <a href="#">The Auth API</a>
            <a href="#">The Report API</a>
            <a href="#">The Analytics API</a>
            <a href="#">The Image API</a>
            <a href="#">More APIs</a>
          </div>
        </div>
        <div class="footer-section">
          <h4>Legals</h4>
          <div class="footer-links">
            <a href="#">Terms & Conditions</a>
            <a href="#">Privacy Policy</a>
          </div>
        </div>
      </div>
    </footer>

    <script src="/static/js/main.js"></script>
  </body>
</html>
