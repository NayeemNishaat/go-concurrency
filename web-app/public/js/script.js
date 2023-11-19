function selectPlan(i, plan, token) {
  Swal.fire({
    title: "Subscribe",
    html: `Subscribe to the ${plan} plan?`,
    showCancelButton: true,
    confirmButtonText: "Subscribe"
  }).then((res) => {
    if (res.isConfirmed) {
      window.location.href = "/subscribe?id=" + i + "&csrfToken=" + token;
    }
  });
}
