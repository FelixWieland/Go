// Obtain the root
const rootElement = document.getElementById('react-root')

class Channel extends React.Component {
	onClick() {
		console.log(this.props.name);
	}
	render() {
		return (
			<li onClick={this.onClick.bind(this)}>{this.props.name}</li>
		)
	}
}

class ChannelList extends React.Component {
	render() {
		return (
			<ul>
				{
					this.props.channels.map(function(channel) {
						return (
							<Channel name={channel.name} />
						)
					})
				}
			</ul>
		)
	}
}

class ChannelForm extends React.Component {
	constructor(props) {
		super(props);
		this.state = {};
	}
	onChange(e) {
		this.setState({
			channelName: e.target.value
		});
	}
	onSubmit(e) {
		var channelName = this.state.channelName;
		this.setState({
			channelName: ''
		});
		this.props.addChannel(channelName);
		e.preventDefault();
	}
	render() {
		return (
			<form onSubmit={this.onSubmit.bind(this)}>
				<input text='text' onChange={this.onChange.bind(this)} value={this.state.channelName}/>
			</form>
		)
	}
}

class ChannelSection extends React.Component {
	constructor(props) {
		super(props);
		this.state = {channels: [
			{name: 'Hardware Support'},
			{name: 'Software Support'}
		]}
	}
	addChannel(name) {
		var {channels} = this.state;
		channels.push({name: name});
		this.setState({
			channels: channels
		})
	}
	render() {
		return (
			<div>
				<ChannelList channels={this.state.channels}/>
				<ChannelForm addChannel={this.addChannel.bind(this)}/>
			</div>
		)
	}
}

ReactDOM.render(
  <ChannelSection/>, 
  rootElement
)
