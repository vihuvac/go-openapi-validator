window.onload = () => {
  const uiElement = document.getElementById('swagger-ui');
  const specURL = uiElement.getAttribute('data-url');
  const absoluteSpecURL = window.location.origin + specURL;

  // Manually set the validator badge as it's often hidden on localhost.
  // NOTE: This will show "ERROR" on localhost because the public validator
  // service (validator.swagger.io) cannot reach your local machine.
  const validatorBase = "https://validator.swagger.io/validator";
  document.getElementById('validator-link').href = validatorBase + "/debug?url=" + encodeURIComponent(absoluteSpecURL);
  document.getElementById('validator-img').src = validatorBase + "?url=" + encodeURIComponent(absoluteSpecURL);

  window.ui = SwaggerUIBundle({
    url: specURL,
    dom_id: '#swagger-ui',
    deepLinking: true,
    presets: [
      SwaggerUIBundle.presets.apis,
      SwaggerUIStandalonePreset
    ],
    plugins: [
      SwaggerUIBundle.plugins.DownloadUrl
    ],
    layout: "StandaloneLayout",
    validatorUrl: validatorBase
  });
};
