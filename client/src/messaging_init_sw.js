import { initializeApp } from "firebase/app";
import { getMessaging, getToken } from "firebase/messaging";

const firebaseConfig = {
  apiKey: "AIzaSyCJzFziE75MvH9cntASm4h5dOPiVO0EBQ8",
  authDomain: "graphical-bus-99503.firebaseapp.com",
  projectId: "graphical-bus-99503",
  storageBucket: "graphical-bus-99503.appspot.com",
  messagingSenderId: "318755619036",
  appId: "1:318755619036:web:d8d3f862ba897ffa073419",
  measurementId: "G-7WD9RH1K6C",
};

function requestPermission() {
  console.log("Requesting permission....");
  Notification.requestPermission().then((permission) => {
    if (permission === "granted") {
      console.log("Notification permission granted");

      const app = initializeApp(firebaseConfig);

      const messaging = getMessaging(app);
      getToken(messaging, {
        vapidKey:
          "BF19CAHw2uZfPzPCba2v6gQ0Mzlm6zoXCeqt3UwF3G_EhMUvPgFv_z6DzNlpQpgze3m3ruJc2GIpuAKTfdvXRvg",
      }).then((token) => {
        if (token) {
          console.log("Firebase Token: ", token);
        } else {
          console.log("Can't get token");
        }
      });
    } else {
      console.log("Do not have permission");
    }
  });
}

requestPermission();
