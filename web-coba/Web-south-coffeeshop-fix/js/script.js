// Toggle class active
const navNav = document.querySelector(".nav-nav");

// ketika hamburger menu diklik
document.querySelector("#hamburger-menu").onclick = () => {
  navNav.classList.toggle("active");
};

// Klik diluar sidebar untuk menghilangkan nav
const hamburger = document.querySelector("#hamburger-menu");

document.addEventListener("click", function (e) {
  if (!hamburger.contains(e.target) && !navNav.contains(e.target)) {
    navNav.classList.remove("active");
  }
});

// TOAST
function showToast(message, duration = 3000) {
  const toastContainer = document.getElementById("toast-container");

  // Create the toast element
  const toast = document.createElement("div");
  toast.classList.add("toast");
  toast.textContent = message;

  // Add the toast to the container
  toastContainer.appendChild(toast);

  // Show the toast
  toast.classList.add("show");

  // Remove the toast after the specified duration
  setTimeout(() => {
    toast.classList.remove("show");
    setTimeout(() => {
      toast.remove();
    }, 300);
  }, duration);
}

// LOGIN
function actionLogin() {
  let emailInput = document.querySelector("#login-email").value;
  let passwordInput = document.querySelector("#login-password").value;

  if (emailInput == "" || passwordInput == "") {
    showToast("Please fill in all the fields");
  } else {
    let data = {
      email: emailInput,
      password: passwordInput,
    };

    fetch("http://localhost:1323/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    })
      .then((response) => {
        if (response.ok) {
          return response.json();
        } else {
          showToast("Wrong email or password");
          throw new Error("Error: " + response.status);
        }
      })
      .then((json) => {
        showToast("Login successful");

        // Success: handle the response data
        console.log("DATA", json.data);
        const token = json.data.access_token;
        const role = json.data.user.access;
        const userId = json.data.user.id;
        const fullname = `${json.data.user.first_name} ${json.data.user.last_name}`;
        const email = json.data.user.email;

        localStorage.setItem("USER_TOKEN", token);
        localStorage.setItem("USER_ROLE", role);
        localStorage.setItem("USER_NAME", fullname);
        localStorage.setItem("USER_EMAIL", email);
        localStorage.setItem("USER_ID", userId);

        // localStorage.setItem("",)

        window.location.replace("index.html");
      })
      .catch((error) => {
        // Error: handle the fetch error
        console.error(error);
      });
  }
}
