<script>

</script>

<div class="root-layout">
  <section class="payment-container">
    <div class="box shadow">
      <img
        src="/assets/icons/payment-success.png"
        alt="Payment Success"
        class="icon"
      />
      <div class="content">
        <h1>Payment Success</h1>
        <p>Lorem ipsum dolor sit amet consectetur, adipisicing elit. Tempore officia sed illo neque quae non blanditiis, voluptate impedit, numquam quis totam iste provident perspiciatis sapiente, et expedita. Doloribus, animi consequuntur.</p>
      </div>
    </div>
  </section>
</div>

<style>
  .root-layout {
    display: block;
		top: 0;
		bottom: 0;
		left: 0;
		right: 0;
		height: 100vh;
		width: 100vw;
    color: var(--gray-009);
  }

  .payment-container {
    display: flex;
    height: 100%;
    width: 100%;
    align-items: center;
    justify-content: center;
    background-color: var(--gray-001);
  }

  .payment-container .box {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    gap: 20px;
    width: 500px;
    height: fit-content;
    background-color: #FFF;
    border-radius: 10px;
    padding: 20px 50px 30px;
  }

  .payment-container .box .icon {
    width: 150px;
    height: auto;
  }

  .payment-container .box .content {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    gap: 10px;
  }

  .payment-container .box .content h1 {
    margin: 0;
    font-size: var(--font-xl);
  }

  .payment-container .box .content p {
    margin: 0;
    text-align: center;
  }

  @media only screen and (max-width : 768px) {
    .payment-container {
      align-items: baseline;
    }

    .payment-container .box {
      width: 100%;
      margin: 20px 10px;
    }
  }
</style>