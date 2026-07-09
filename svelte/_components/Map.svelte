<script>
  import { onMount, onDestroy } from 'svelte';
  import 'maplibre-gl/dist/maplibre-gl.css';
	import { Map, NavigationControl, Marker } from 'maplibre-gl';

  let map;          // Map object from lib
  let mapContainer; // Map container DOM

  const apiKey = 'tOiu6Qb0Q3hToL89vQ4u';

  const initialLat = -2.548926;
  const initialLng = 118.0148634;

  // +=========== EXPORTED VARIABLES ===========+
  export let Coord = {
    lng: initialLng,
    lat: initialLat,
    zoom: 4
  };

  export let IsClickable = false;
  export let IsDraggable = false;

  /** @param {[number, number]} lngLat */
  export let OnClickMap = function(lngLat) {};

  /** @param {[number, number]} location */
  export function setCenter(location) {
    map.setZoom(Coord.zoom);
    console.log('Location: ', location);
    map.setCenter(location);
  }

  // +===========================================+

  function isValidCoordinates(lng, lat) {
    return (
      typeof lng === 'number' &&
      typeof lat === 'number' &&
      !isNaN(lng) &&
      !isNaN(lat) &&
      lng >= -180 && lng <= 180 &&
      lat >= -90 && lat <= 90
    );
  }

  let marker = new Marker({
    color: '#1877F2',
    draggable: IsDraggable
  });

  onMount(() => {
    if (!isValidCoordinates(Coord.lng, Coord.lat)) {
      Coord.lng = initialLng;
      Coord.lat = initialLat;
    }
    map = new Map({
      container: mapContainer,
      style: `https://api.maptiler.com/maps/streets/style.json?key=${apiKey}`,
      center: [Coord.lng, Coord.lat],
      zoom: Coord.zoom
    });
    map.addControl(new NavigationControl(), 'top-right');
    map.on('load', () => {
      marker.setLngLat(Coord).addTo(map);
      if (IsClickable) {
        map.on('click', (/** @type any */ e) => {
          /** @type {import('maplibre-gl').LngLatLike} */
          const lngLat = [e.lngLat.lng, e.lngLat.lat];

          OnClickMap(lngLat);
          marker.setLngLat(lngLat).addTo(map);
        });
      }
    });

    if (IsDraggable) {
      marker.on('dragend', (/** @type any */ e) => {
        /** @type {import('maplibre-gl').LngLatLike} */
        const lngLat = [e.target._lngLat.lng, e.target._lngLat.lat]; 

        marker.setLngLat(lngLat).addTo(map);
        map.setCenter(lngLat);
        OnClickMap(lngLat);
      })
    }
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