<script>
  import { onMount, onDestroy } from 'svelte';
  import 'maplibre-gl/dist/maplibre-gl.css';
	import { Map, NavigationControl, Marker } from 'maplibre-gl';

  let map;          // Map object from lib
  let mapContainer; // Map container DOM

  const apiKey = 'tOiu6Qb0Q3hToL89vQ4u';

  export let coord = {
    lng: 118.0148634,
    lat: -2.548926,
    zoom: 4
  };

  /** @param {[number, number]} location */
  export function setCenter(location) {
    map.setZoom(11);
    map.setCenter(location);
  }

  onMount(() => {
    map = new Map({
      container: mapContainer,
      style: `https://api.maptiler.com/maps/streets/style.json?key=${apiKey}`,
      center: [coord.lng, coord.lat],
      zoom: coord.zoom
    });
    map.addControl(new NavigationControl(), 'top-right');
    new Marker({color: '#1877F2'}).setLngLat(coord).addTo(map);
  });

  onDestroy( () => map.remove() );
</script>

<div class="map-wrapper">
  <div
    class="map"
    id="map"
    bind:this={mapContainer}
  ></div>
</div>

<style>
  :global(.maplibregl-popup-close-button) {
    color: var(--red-005) !important;
    font-weight: 600;
  }
  .map-wrapper {
    position: relative;
    width: 100%;
    height: 100%;
    overflow: hidden;
    border-radius: 10px;
    border: 1px solid var(--gray-003);
  }
  
  .map {
    position: absolute;
    width: 100%;
    height: 100%;
  }
</style>