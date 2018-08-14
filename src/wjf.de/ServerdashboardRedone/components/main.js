
// Obtain the root
const rootElement = document.getElementById('root')
// Create a ES6 class component

class List extends React.Component {
  render() {
      return (
        <div className="container well">
          <div className="row">
          <h1>Test for: {this.props.name}</h1>
            <ul>
              <li>Test 1</li>
              <li>Test 2</li>
              <li>Test 2</li>
            </ul>
          </div>
        </div>
      );
    }
}



// Create a function to wrap up your component
function App(){
  return(
    <div>
      <Menu />
    </div>
  )
}


// Use the ReactDOM.render to show your component on the browser
ReactDOM.render(
  <App />,
  rootElement
)
