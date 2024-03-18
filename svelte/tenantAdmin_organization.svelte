<script>
  import MainLayout from './_layouts/mainLayout.svelte';
  import { onMount } from 'svelte';
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import RiBuildingsCommunityLine from 'svelte-icons-pack/ri/RiBuildingsCommunityLine';
  import RiSystemAddBoxLine from 'svelte-icons-pack/ri/RiSystemAddBoxLine';
  import RiDesignPencilLine from 'svelte-icons-pack/ri/RiDesignPencilLine';
  import RiSystemDeleteBinLine from 'svelte-icons-pack/ri/RiSystemDeleteBinLine';
  import RiSystemInformationLine from 'svelte-icons-pack/ri/RiSystemInformationLine';
  import RiBuildingsBuilding2Line from 'svelte-icons-pack/ri/RiBuildingsBuilding2Line';
  import RiUserTeamLine from "svelte-icons-pack/ri/RiUserTeamLine";
  import RiBusinessBriefcaseLine from "svelte-icons-pack/ri/RiBusinessBriefcaseLine";

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
    <div class="org">
      <div class="company">
        <div class="info">
          <div class="label">
            <Icon
              color="#FFF"
              className="icon"
              size="16"
              src={RiBuildingsCommunityLine}
            />
          </div>
          <span class="title">{orgs[0].name}</span>
        </div>
        <div class="options">
          <button class="btn" title="Add department">
            <Icon
              color="var(--gray-006)"
              className="icon"
              size="17"
              src={RiSystemAddBoxLine}
            />
          </button>
          <button class="btn" title="Edit company">
            <Icon
              color="var(--gray-006)"
              className="icon"
              size="17"
              src={RiDesignPencilLine}
            />
          </button>
          <button class="btn" title="Delete company">
            <Icon
              color="var(--gray-006)"
              className="icon"
              size="17"
              src={RiSystemDeleteBinLine}
            />
          </button>
          <button class="btn" title="Info">
            <Icon
              color="var(--gray-006)"
              className="icon"
              size="17"
              src={RiSystemInformationLine}
            />
          </button>
        </div>
      </div>
      <div class="department">
        <div class="info">
          <span class="h-line"></span>
          <div class="label">
            <Icon
              color="#FFF"
              className="icon"
              size="16"
              src={RiBuildingsBuilding2Line}
            />
          </div>
          <span class="title">{orgs[0].name}</span>
        </div>
        <div class="options">
          <button class="btn" title="Add department">
            <Icon
              color="var(--gray-006)"
              className="icon"
              size="17"
              src={RiSystemAddBoxLine}
            />
          </button>
          <button class="btn" title="Edit company">
            <Icon
              color="var(--gray-006)"
              className="icon"
              size="17"
              src={RiDesignPencilLine}
            />
          </button>
          <button class="btn" title="Delete company">
            <Icon
              color="var(--gray-006)"
              className="icon"
              size="17"
              src={RiSystemDeleteBinLine}
            />
          </button>
          <button class="btn" title="Info">
            <Icon
              color="var(--gray-006)"
              className="icon"
              size="17"
              src={RiSystemInformationLine}
            />
          </button>
        </div>
      </div>
      <div class="division">
        <div class="info">
          <span class="h-line"></span>
          <div class="label">
            <Icon
              color="#FFF"
              className="icon"
              size="16"
              src={RiUserTeamLine}
            />
          </div>
          <span class="title">{orgs[0].name}</span>
        </div>
        <div class="options">
          <button class="btn" title="Add department">
            <Icon
              color="var(--gray-006)"
              className="icon"
              size="17"
              src={RiSystemAddBoxLine}
            />
          </button>
          <button class="btn" title="Edit company">
            <Icon
              color="var(--gray-006)"
              className="icon"
              size="17"
              src={RiDesignPencilLine}
            />
          </button>
          <button class="btn" title="Delete company">
            <Icon
              color="var(--gray-006)"
              className="icon"
              size="17"
              src={RiSystemDeleteBinLine}
            />
          </button>
          <button class="btn" title="Info">
            <Icon
              color="var(--gray-006)"
              className="icon"
              size="17"
              src={RiSystemInformationLine}
            />
          </button>
        </div>
      </div>
      <div class="division">
        <div class="info">
          <span class="h-line"></span>
          <div class="label">
            <Icon
              color="#FFF"
              className="icon"
              size="16"
              src={RiUserTeamLine}
            />
          </div>
          <span class="title">{orgs[0].name}</span>
        </div>
        <div class="options">
          <button class="btn" title="Add department">
            <Icon
              color="var(--gray-006)"
              className="icon"
              size="17"
              src={RiSystemAddBoxLine}
            />
          </button>
          <button class="btn" title="Edit company">
            <Icon
              color="var(--gray-006)"
              className="icon"
              size="17"
              src={RiDesignPencilLine}
            />
          </button>
          <button class="btn" title="Delete company">
            <Icon
              color="var(--gray-006)"
              className="icon"
              size="17"
              src={RiSystemDeleteBinLine}
            />
          </button>
          <button class="btn" title="Info">
            <Icon
              color="var(--gray-006)"
              className="icon"
              size="17"
              src={RiSystemInformationLine}
            />
          </button>
        </div>
      </div>
      <div class="job">
        <div class="info">
          <span class="h-line"></span>
          <div class="label">
            <Icon
              color="#FFF"
              className="icon"
              size="16"
              src={RiBusinessBriefcaseLine}
            />
          </div>
          <span class="title">{orgs[0].name}</span>
        </div>
        <div class="options">
          <button class="btn" title="Add department">
            <Icon
              color="var(--gray-006)"
              className="icon"
              size="17"
              src={RiSystemAddBoxLine}
            />
          </button>
          <button class="btn" title="Edit company">
            <Icon
              color="var(--gray-006)"
              className="icon"
              size="17"
              src={RiDesignPencilLine}
            />
          </button>
          <button class="btn" title="Delete company">
            <Icon
              color="var(--gray-006)"
              className="icon"
              size="17"
              src={RiSystemDeleteBinLine}
            />
          </button>
          <button class="btn" title="Info">
            <Icon
              color="var(--gray-006)"
              className="icon"
              size="17"
              src={RiSystemInformationLine}
            />
          </button>
        </div>
      </div>
      <div class="department">
        <div class="info">
          <span class="h-line"></span>
          <div class="label">
            <Icon
              color="#FFF"
              className="icon"
              size="16"
              src={RiBuildingsBuilding2Line}
            />
          </div>
          <span class="title">{orgs[0].name}</span>
        </div>
        <div class="options">
          <button class="btn" title="Add department">
            <Icon
              color="var(--gray-006)"
              className="icon"
              size="17"
              src={RiSystemAddBoxLine}
            />
          </button>
          <button class="btn" title="Edit company">
            <Icon
              color="var(--gray-006)"
              className="icon"
              size="17"
              src={RiDesignPencilLine}
            />
          </button>
          <button class="btn" title="Delete company">
            <Icon
              color="var(--gray-006)"
              className="icon"
              size="17"
              src={RiSystemDeleteBinLine}
            />
          </button>
          <button class="btn" title="Info">
            <Icon
              color="var(--gray-006)"
              className="icon"
              size="17"
              src={RiSystemInformationLine}
            />
          </button>
        </div>
      </div>
      <div class="division">
        <div class="info">
          <span class="h-line"></span>
          <div class="label">
            <Icon
              color="#FFF"
              className="icon"
              size="16"
              src={RiUserTeamLine}
            />
          </div>
          <span class="title">{orgs[0].name}</span>
        </div>
        <div class="options">
          <button class="btn" title="Add department">
            <Icon
              color="var(--gray-006)"
              className="icon"
              size="17"
              src={RiSystemAddBoxLine}
            />
          </button>
          <button class="btn" title="Edit company">
            <Icon
              color="var(--gray-006)"
              className="icon"
              size="17"
              src={RiDesignPencilLine}
            />
          </button>
          <button class="btn" title="Delete company">
            <Icon
              color="var(--gray-006)"
              className="icon"
              size="17"
              src={RiSystemDeleteBinLine}
            />
          </button>
          <button class="btn" title="Info">
            <Icon
              color="var(--gray-006)"
              className="icon"
              size="17"
              src={RiSystemInformationLine}
            />
          </button>
        </div>
      </div>
      <div class="department">
        <div class="info">
          <span class="h-line"></span>
          <div class="label">
            <Icon
              color="#FFF"
              className="icon"
              size="16"
              src={RiBuildingsBuilding2Line}
            />
          </div>
          <span class="title">{orgs[0].name}</span>
        </div>
        <div class="options">
          <button class="btn" title="Add department">
            <Icon
              color="var(--gray-006)"
              className="icon"
              size="17"
              src={RiSystemAddBoxLine}
            />
          </button>
          <button class="btn" title="Edit company">
            <Icon
              color="var(--gray-006)"
              className="icon"
              size="17"
              src={RiDesignPencilLine}
            />
          </button>
          <button class="btn" title="Delete company">
            <Icon
              color="var(--gray-006)"
              className="icon"
              size="17"
              src={RiSystemDeleteBinLine}
            />
          </button>
          <button class="btn" title="Info">
            <Icon
              color="var(--gray-006)"
              className="icon"
              size="17"
              src={RiSystemInformationLine}
            />
          </button>
        </div>
      </div>
      <div class="company">
        <div class="info">
          <div class="label">
            <Icon
              color="#FFF"
              className="icon"
              size="16"
              src={RiBuildingsCommunityLine}
            />
          </div>
          <span class="title">{orgs[0].name}</span>
        </div>
        <div class="options">
          <button class="btn" title="Add department">
            <Icon
              color="var(--gray-006)"
              className="icon"
              size="17"
              src={RiSystemAddBoxLine}
            />
          </button>
          <button class="btn" title="Edit company">
            <Icon
              color="var(--gray-006)"
              className="icon"
              size="17"
              src={RiDesignPencilLine}
            />
          </button>
          <button class="btn" title="Delete company">
            <Icon
              color="var(--gray-006)"
              className="icon"
              size="17"
              src={RiSystemDeleteBinLine}
            />
          </button>
          <button class="btn" title="Info">
            <Icon
              color="var(--gray-006)"
              className="icon"
              size="17"
              src={RiSystemInformationLine}
            />
          </button>
        </div>
      </div>
    </div>
  </div>
</MainLayout>

<style>
  .org {
    display: flex;
    flex-direction: column;
    gap: 10px;
  }

  .org .company,
  .org .department,
  .org .division,
  .org .job {
    display: flex;
    align-items: center;
    flex-direction: row;
    gap: 40px;
    padding: 5px 0;
  }

  .org .department {
    padding-left: 30px;
  }

  .org .department .info .label {
    background-color: var(--orange-006);
  }

  .org .division {
    padding-left: 50px;
  }

  .org .division .info .label {
    background-color: var(--green-006);
  }

  .org .job {
    padding-left: 70px;
  }

  .org .job .info .label {
    background-color: var(--gray-006);
  }

  .org .info {
    display: flex;
    align-items: center;
    flex-direction: row;
    gap: 10px;
    position: relative;
  }

  .org .info .h-line {
    position: absolute;
    left: -15px;
    width: 1px;
    height: 50px;
    background-color: var(--gray-003);
  }

  .org .info .label {
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

  .org .info .title {
    font-size: 18px;
    font-weight: 600;
  }

  .org .company:hover .options,
  .org .department:hover .options,
  .org .division:hover .options,
  .org .job:hover .options {
    display: flex;
    cursor: pointer;
  }

  .org .options {
    display: none;
    flex-direction: row;
    align-items: center;
    gap: 5px;
  }

  .org .options .btn {
    display: flex;
    justify-content: center;
    align-items: center;
    border: none;
    background-color: transparent;
    border-radius: 5px;
    padding: 5px;
    cursor: pointer;
  }

  .org .options .btn:hover {
    background-color: var(--gray-002);
  }

  .org .options .btn:active {
    background-color: var(--gray-003);
  }
</style>