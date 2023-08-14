<script>
	import FlightCard from './flight_card.svelte';
	export let data;
	const { flight } = data;
</script>

<main>
	<head>
		<title>Find A Flight</title>
	</head>

	<div class="card-container">
		{#await flight}
			<p>loading...</p>
		{:then flight}
			{#each flight as { Departure_time, Carrier_iata, Flight_Number, Source_Airport, Destination_Airport, Destination_Country, Destination_City, Best_Flight_Price, Best_Flight }}
				<FlightCard
					{Departure_time}
					{Carrier_iata}
					{Flight_Number}
					{Source_Airport}
					{Destination_Airport}
					{Destination_Country}
					{Destination_City}
					{Best_Flight_Price}
					{Best_Flight}
				/>
			{/each}
		{:catch error}
			<p>{error.message}</p>
		{/await}
	</div>
</main>

<style>
	.card-container {
		display: flex;
		flex-wrap: wrap;
		justify-content: space-evenly;
	}
</style>
