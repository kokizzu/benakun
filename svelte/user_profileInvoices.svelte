<script>
  /** @typedef {import('./_components/types/user').User} User */
  /** @typedef {import('./_components/types/access').Access} Access */
  /** @typedef {import('./_components/types/invoicePayment').InvoicePayment} InvoicePayment */

  import { Icon } from './node_modules/svelte-icons-pack/dist';
  import { IoClose } from './node_modules/svelte-icons-pack/dist/io';
  import { datetime, localeDatetime } from './_components/xformatter.js';
  import InputBox from './_components/InputBox.svelte';
  import SubmitButton from './_components/SubmitButton.svelte';
  import MainLayout from './_layouts/mainLayout.svelte';
  import ProfileSubMenu from './_components/partials/ProfileSubMenu.svelte';

  let user      = /** @type {User} */ ({/* user */});
  let segments  = /** @type {Access} */ ({/* segments */});
  let invoices  = /** @type {InvoicePayment[]} */ ([/* invoices */]);

  console.log('User:', user);
  console.log('Segments:', segments);
  console.log('Invoices:', invoices);
</script>

<MainLayout>
  <ProfileSubMenu />
  <div class="invoices-container">
    <p>Invoices here</p>
  </div>
</MainLayout>


<style>
  
</style>
