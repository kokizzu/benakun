<script>
  import MainLayout from './_layouts/mainLayout.svelte';
  import { onMount } from 'svelte';
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import RiBuildingsCommunityLine from 'svelte-icons-pack/ri/RiBuildingsCommunityLine';

  let segments = {/* segments */};
  let user = {/* user */};

  /**
    * @typedef {Object} Org
    * @property {string} id
    * @property {string} name
    * @property {number} orgType
    * @property {string} parentId
    * @property {string} tenantCode
    * @property {number} createdAt
    * @property {string} createdBy
    * @property {number} updatedAt
    * @property {string} updatedBy
    * @property {number} deletedAt
    * @property {Array<Org>} children
    */
  /**
   * @type {Array<Org>}
   */
  let orgs = [/* orgs */];

  onMount(() => {
    console.log('Segments:', segments);
    console.log('User:', user);
    console.log('Orgs:', orgs);
  })
</script>

<MainLayout>
  <div>
    {#if orgs && orgs.length}
      {#each orgs as org, _ (org.id)}
        <div class="org">
          <div class="label">
            <Icon
              color="#FFF"
              className="icon"
              size="16"
              src={RiBuildingsCommunityLine}
            />
          </div>
          <span class="text">{org.name}</span>
        </div>
      {/each}
    {/if}
  </div>
</MainLayout>

<style>
  .org {
    display: flex;
    align-items: center;
    flex-direction: row;
    gap: 10px;
  }

  .org > .label {
    display: flex;
    align-items: center;
    justify-content: center;
    width: fit-content;
    height: fit-content;
    padding: 7px;
    border-radius: 9999px;
    background-color: var(--blue-006);
    color: #FFF;
  }

  .org > .text {
    font-size: 18px;
    font-weight: 600;
  }
</style> 