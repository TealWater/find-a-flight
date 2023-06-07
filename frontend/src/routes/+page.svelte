<script>
	import FlightCard from './flight_card.svelte';
	var flights;
	const fetchData = async () => {
		const res = await fetch('https://find-a-flight-backend.onrender.com/getData');
		const data = await res.json();
		console.log(data);
		return data;
	};
	console.log('hi mom!');
</script>

<main>
	<head>
		<title>Find A Flight</title>
	</head>


	<div class="card-container">
		{#await fetchData()}
			<p>loading...</p>
		{:then data}
			{#each data as { Departure_time, Carrier_iata, Flight_Number, Source_Airport, Destination_Airport, Destination_Country, Destination_City, Best_Flight_Price, Best_Flight }}
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
