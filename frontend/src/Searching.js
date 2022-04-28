
import React from 'react'
import Header from './component/Header'
import Logo from './component/Logo'
import SearchBar from './component/SearchBar'
const Searching = () => {
  return (
    <div>
        <Header/>
        <div className="otherpages">
            <Logo/>
            <SearchBar data={'1 April 2022 - Mhs IF - HIV - 75% - False'} data2 = {'yabai'}/>
        </div>
    </div>
  )
}
export default Searching

/*
import React, { Component } from 'react'
import Header from './component/Header'
import Logo from './component/Logo'
import SearchBar from './component/SearchBar'
import axios from 'axios'
import Output from './component/Output'

export class Searching extends Component {
  constructor(props) {
    super(props)
  
    this.state = {
       posts: []
    }
  }

  componentDidMount(){
    axios.get('https://jsonplaceholder.typicode.com/posts')
    .then(response => {
      console.log(response)
      this.setState({ posts: response.data})
    })
    .catch(error => {
      console.log(error)
    })
  }

  render() {
    const { posts } = this.state
    return (
      <div>
        <Header/>
        <div className="otherpages">
            <Logo/>
            <SearchBar/>
            {
              posts.length ?
              posts.map(post => <Output text={post.id} text2 = {post.title}/>) : null
            }
        </div>
      </div>
    )
  }
}

export default Searching

*/