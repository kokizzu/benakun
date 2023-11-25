<script>
  const year = new Date().getFullYear();
</script>

<footer>
  <p>Copyright © {year}</p>
  <nav>
    <a href="/privacy">Privacy</a>
    <a href="/tos">Terms and Conditions </a>
    <a href="/about">About Us</a>
  </nav>
</footer>

<style>
  footer {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    margin: 30px 40px 0 40px;
    padding: 20px 0;
    border-top: 1px solid var(--gray-003);
    color: var(--gray-007);
  }
  footer p {
    margin: 0;
  }
  footer nav {
    display: flex;
    flex-direction: row;
  }
  footer nav a {
    text-decoration: none;
    margin-left: 15px;
    color: var(--gray-007);
  }
  footer nav a:hover {
    color: var(--purple-002);
  }
</style>
