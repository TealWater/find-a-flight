<script>
	import { airports } from './load_iata.js';
	import Airport from './airports.svelte';
	import { goto } from '$app/navigation';

	/* FILTERING countres DATA BASED ON INPUT */
	let filteredAirports = [];
	// $: console.log(filteredAirports)

	const filterAirports = () => {
		let storageArr = [];
		if (inputValue) {
			airports.forEach((airport) => {
				if (airport.toLowerCase().startsWith(inputValue.toLowerCase())) {
					storageArr = [...storageArr, makeMatchBold(airport)];
				}
			});
		}
		filteredAirports = storageArr;
	};

	/* HANDLING THE INPUT */
	let searchInput; // use with bind:this to focus element
	let inputValue = '';

	$: if (!inputValue) {
		filteredAirports = [];
		hiLiteIndex = null;
	}

	const clearInput = () => {
		inputValue = '';
		searchInput.focus();
	};

	const setInputVal = (airportName) => {
		inputValue = removeBold(airportName);
		filteredAirports = [];
		hiLiteIndex = null;
		document.querySelector('#airport-input').focus();
	};

	const submitValue = () => {
		if (inputValue) {
			console.log(`${inputValue} is submitted!`);
			goto(`/booking/${inputValue}`);
			setTimeout(clearInput, 1000);
		} else {
			alert("You didn't type anything.");
		}
	};

	const makeMatchBold = (str) => {
		// replace part of (airport name === inputValue) with strong tags
		let matched = str.substring(0, inputValue.length);
		let makeBold = `<strong>${matched}</strong>`;
		let boldedMatch = str.replace(matched, makeBold);
		return boldedMatch;
	};

	const removeBold = (str) => {
		//replace < and > all characters between
		return str.replace(/<(.)*?>/g, '');
		// return str.replace(/<(strong)>/g, "").replace(/<\/(strong)>/g, "");
	};

	/* NAVIGATING OVER THE LIST OF AIRPORTS W HIGHLIGHTING */
	let hiLiteIndex = null;
	//$: console.log(hiLiteIndex);
	$: hiLitedAirport = filteredAirports[hiLiteIndex];

	const navigateList = (e) => {
		if (e.key === 'ArrowDown' && hiLiteIndex <= filteredAirports.length - 1) {
			hiLiteIndex === null ? (hiLiteIndex = 0) : (hiLiteIndex += 1);
		} else if (e.key === 'ArrowUp' && hiLiteIndex !== null) {
			hiLiteIndex === 0 ? (hiLiteIndex = filteredAirports.length - 1) : (hiLiteIndex -= 1);
		} else if (e.key === 'Enter') {
			setInputVal(filteredAirports[hiLiteIndex]);
		} else {
			return;
		}
	};
</script>

<svelte:window on:keydown={navigateList} />

<center>
	<form autocomplete="off" on:submit|preventDefault={submitValue}>
		<div class="autocomplete">
			<input
				id="airport-input"
				type="text"
				placeholder="Search Airport Codes (JFK)"
				bind:this={searchInput}
				bind:value={inputValue}
				on:input={filterAirports}
			/>
		</div>

		<input type="submit" />

		<!-- FILTERED LIST OF AIRPORTS -->
		{#if filteredAirports.length > 0}
			<ul id="autocomplete-items-list">
				{#each filteredAirports as airport, i}
					<Airport
						itemLabel={airport}
						highlighted={i === hiLiteIndex}
						on:click={() => setInputVal(airport)}
					/>
				{/each}
			</ul>
		{/if}
	</form>
</center>

<style>
	div.autocomplete {
		/*the container must be positioned relative:*/
		position: relative;
		display: inline-block;
		width: 300px;
	}
	input {
		border: 1px solid transparent;
		background-color: #f1f1f1;
		padding: 10px;
		font-size: 16px;
		top: 38em;
		bottom: 38em;
	}
	input[type='text'] {
		position: relative;
		background-color: #f1f1f1;
		width: 100%;
	}
	input[type='submit'] {
		position: relative;
		margin-left: 1.3em;
		background-color: DodgerBlue;
		color: #fff;
	}

	#autocomplete-items-list {
		position: relative;
		/*margin: 0;*/
		padding: 0;
		/*top: 0;*/
		top: 37em;
		bottom: 37em;
		right: 2em;
		width: 12.6%;
		border: 1px solid #ddd;
		background-color: #ddd;
	}
</style>
