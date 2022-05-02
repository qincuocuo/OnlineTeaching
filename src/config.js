const getLocationOrigin = () => {
  return (
    window.location.protocol +
    "//" +
    window.location.hostname +
    (window.location.port ? ":" + window.location.port : "")
  );
};

const version = "V1.0";

export default {
  version,
  getLocationOrigin
};
