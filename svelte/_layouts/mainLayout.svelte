<script>
  /** @typedef {import('../_components/types/access').Access} Access */
  /** @typedef {import('../_components/types/user').User} User*/

  import SideMenu from '../_components/partials/SideMenu.svelte';
  import Navbar from '../_components/partials/Navbar.svelte';
  import Footer from '../_components/partials/Footer.svelte';
  import { IsShrinkMenu } from '../_components/uiState';
  import { onMount } from 'svelte';
  
  let segments = /** @type Access */ ({/* segments */});
  let user = /** @type User */ ({/* user */});

  let viewportWidth = window.innerWidth;

  onMount(() => {
    if (viewportWidth > 768) {
      const isShrink = localStorage.getItem('IsShrinkMenu');
      IsShrinkMenu.set(JSON.parse(isShrink));
    } else {
      localStorage.setItem('IsShrinkMenu', 'true');
      IsShrinkMenu.set(true);
      document.addEventListener('click', (e) => {
        const elm = /** @type HTMLElement */ (e.target);
        if (elm.tagName == 'a') {
          console.log('TRIGGERED')
          $IsShrinkMenu = true;
        }
      })
    }
  })
</script>

<div class="root_layout">
  <div class="root_container">
    <SideMenu access={segments} />
    <div class="root_content {$IsShrinkMenu ? 'shrink' : 'expand'}">
      <Navbar {user} />
      <div class="content">
        <main><slot /></main>
        <Footer />
      </div>
    </div>
  </div>
</div>

<style>
  .root_layout {
    display: block;
		top: 0;
		bottom: 0;
		left: 0;
		right: 0;
		height: 100vh;
		width: 100vw;
  }

  .root_layout .root_container {
    height: 100%;
		width: 100%;
		display: flex;
  }

  .root_layout .root_container .root_content {
		display: flex;
		flex-direction: column;
		-webkit-box-orient: vertical;
		-webkit-box-direction: normal;
		overflow-y: auto;
		min-height: calc(100vh - var(--navbar-height));
		transition: 0.3s;
		width: 100%;
  }

  .root_layout .root_container .root_content.shrink {
    margin-left: var(--sidemenu-width-sm);
  }

  .root_layout .root_container .root_content.expand {
		margin-left: var(--sidemenu-width);
  }

  .root_layout .root_container .root_content .content {
    width: 100%;
    padding: 20px;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    gap: 30px;
    min-height: calc(100vh - var(--navbar-height));
    max-height: fit-content;
    overflow: inherit;
  }

  @media only screen and (max-width : 768px) {
    .root_layout .root_container .root_content.shrink {
      margin-left: 0;
    }

    .root_layout .root_container .root_content.expand {
      margin-left: 0;
    }

    .root_layout .root_container .root_content .content {
      padding: 15px;
    }
  }
</style>