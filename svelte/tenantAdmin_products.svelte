<script>
  import MainLayout from './_layouts/mainLayout.svelte';

  /** @typedef {import('./_components/types/master.js').Field} Field */
	/** @typedef {import('./_components/types/access.js').Access} Access */
  /** @typedef {import('./_components/types/master.js').PagerIn} PagerIn */
	/** @typedef {import('./_components/types/master.js').PagerOut} PagerOut */

  let segments = /** @type Access */ ({/* segments */});
  let fields = /** @type Field[] */ ([/* fields */]);
  let pager = /** @type PagerOut */ ({/* pager */});
  let product = {/* product */}
  let products = {/* products */}

  console.log('Segments:', segments);
  console.log('Fields:', fields);
  console.log('Pager:', pager);
  console.log('Product:', product);
  console.log('Products:', products);
</script>

<MainLayout>
  <div>
    <h2>Products</h2>
  </div>
</MainLayout>

<style>
</style>
