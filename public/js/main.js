((w) => {

  const init = () => {
    console.log('ready!')
  }

  w.onload = () => {
    if (document.readyState === 'complete') {
      init()
    }
  }

})(window)
