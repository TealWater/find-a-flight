/** @type {import('./$types').PageServerLoad} */
export async function load({ fetch, params }) {
    const fetchData = async (airport) => {
        const res = await fetch(`https://find-a-flight-backend.onrender.com/search?airport=${airport}`);
        const data = await res.json();
        return data;
    };

    return {
        flight: fetchData(params.slug)
    }
};