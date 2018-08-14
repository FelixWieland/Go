class Menu extends React.Component {
    constructor(props) {
      super(props);
      this.state = { active: false }
    }
    render() {
      return (
        <div>
          <button className={this.state.active && 'active'}
            onClick={ () => this.setState({active: !this.state.active}) }>Click me</button>
        </div>
      )
    }
}