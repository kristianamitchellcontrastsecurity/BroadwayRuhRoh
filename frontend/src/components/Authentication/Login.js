import React, { Component } from 'react';
import { Button, Input, Divider, Segment, List, Grid } from "semantic-ui-react";

export default class Login extends Component {

  render() {
    return(
      <Grid style={{height: '100vh'}} textAlign='center' verticalAlign='middle' columns={3}>
      <Grid.Row>
          {/* Column 1 */}
          <Grid.Column>
          </Grid.Column>
          {/* Column 2 */}
          <Grid.Column>
              <Segment padded>
              <h1>Yalp</h1>
              <Divider horizontal>Sign In</Divider>
              <List>
                  <List.Item>
                      <Input name='email' icon='user' iconPosition='left' placeholder='Enter email...'/>
                      {/* <Input name='email' icon='user' iconPosition='left' placeholder='Enter email...' onChange={this.handleChange}/> */}
                  </List.Item>
                  <List.Item>
                      <Input name='password' icon='lock' iconPosition='left' placeholder='Enter password...' type="password"/>
                      {/* <Input name='password' icon='lock' iconPosition='left' placeholder='Enter password...' type="password" onChange={this.handleChange} /> */}
                  </List.Item>
                  <Button onClick={this.handleClick}>Login</Button>
                  {/* <Button onClick={this.handleClick}>Login</Button> */}
              </List>
              <Divider horizontal>OR</Divider>
              {/* <StyledFirebaseAuth uiConfig={uiConfig} firebaseAuth={firebase.auth()}/> */}
              </Segment>
          </Grid.Column>
          {/* Column 3 */}
          <Grid.Column>
          </Grid.Column>
      </Grid.Row>
  </Grid>
    );
  }

};