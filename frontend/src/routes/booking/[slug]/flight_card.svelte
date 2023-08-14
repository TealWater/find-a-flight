<script>
	export let Departure_time = '';
	export let Carrier_iata = '';
	export let Flight_Number = '';
	export let Source_Airport = '';
	export let Destination_Airport = '';
	export let Destination_Country = '';
	export let Destination_City = '';
	export let Best_Flight_Price = '';
	export let Best_Flight = '';
</script>

<article class="contact-card">
	<h2>
		<slot name="destination">
			<span class="missing">Destination: {Destination_Airport}</span>
		</slot>
	</h2>

	<div class="carrier">
		<slot name="Carrier">
			<span class="missing">Airline: {Carrier_iata}</span>
		</slot>
	</div>

	<div class="flight_num">
		<slot name="flight_num">
			<span class="missing">Flight Number: {Flight_Number}</span>
		</slot>
	</div>

	<div class="source_airport">
		<slot name="_source_airport">
			<span class="missing">Departure Airport: {Source_Airport}</span>
		</slot>
	</div>

	<div class="time">
		<slot name="time">
			<span class="missing">Departure Date: {Departure_time}</span>
		</slot>
	</div>

	<div class="destination_country">
		<slot name="destination_country">
			<span class="missing">Destination Country: {Destination_Country}</span>
		</slot>
	</div>

	<div class="destination_city">
		<slot name="destination_city">
			<span class="missing">Destination City: {Destination_City}</span>
		</slot>
	</div>

	<div class="best_price">
		<slot name="best_price">
			<a href={Best_Flight}>
				<span class="missing">Best Price: {Best_Flight_Price}</span>
			</a>
		</slot>
	</div>
</article>

<style>
	.contact-card {
		width: 300px;
		border: 1px solid #aaa;
		border-radius: 2px;
		box-shadow: 2px 2px 8px rgba(0, 0, 0, 0.1);
		padding: 1em;
	}

	h2 {
		padding: 0 0 0.2em 0;
		margin: 0 0 1em 0;
		border-bottom: 1px solid #ff3e00;
	}

	.carrier,
	.flight_num,
	.time,
	.best_price,
	.source_airport,
	.destination_country,
	.destination_city {
		padding: 0 0 0 1.5em;
		background: 0 50% no-repeat;
		background-size: 1em 1em;
		margin: 0 0 0.5em 0;
		line-height: 1.2;
	}
	.missing {
		color: #999;
	}
</style>
