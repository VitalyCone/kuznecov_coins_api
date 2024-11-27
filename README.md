<h1>Kuznecov Coins API<h1>
<h5>
</p>Usage</p>
<p>1. Clone this repository</p>
<code>git clone https://github.com/VitalyCone/kuznecov_coins_api</code>
<p></p>
<p>1. Build and run <i>docker-compose.yml</i></p>
<code>docker-compose up --build kzcv-coins</code>
<p></p>
<p>2. Apply migrations</p>
<code>migrate -path migrations -database "postgres://admin:admin@localhost:5002/kuznecov_coins?sslmode=disable" up</code>
<p></p>
<p>3. Join in swagger on your browser for check API documentation</p>
<a href="http://localhost:8002/swagger/index.html">http://localhost:8002/swagger/index.html</a>
</h5>
