# User Manual

- **Project Title**: PlusVasis
- **Student 1:** Jason Henderson, 19309916
- **Student 2:** Conor Joyce, 19425804
- **Supervisor:** Dr Stephen Blott
- **Date Completed:** 05/05/23

## Table of Contents

[TOC]

## The Web App

### Setup

To get started with our web application, simply visit the following domain:

https://app.plusvasis.xyz

Our application is fully hosted on the PlusVasis production environment, which is accessible from anywhere with an internet connection. This means you can use our application from any device or location that has an internet connection. No additional software installation or setup is required, just visit the above domain and you're ready to go.

### Getting Started

Once you arrive, you'll see our login screen, where you can either sign in to an existing account or create a new one. You have two options for signing in: you can either use your email address and password, or sign in via GitHub.

If you're new to PlusVasis and need to create an account, simply select the 'Create Account' option on the login screen and follow the prompts to set up your account. If you already have an account, enter your login credentials and you'll be taken to your dashboard.

That's all there is to it! Once you're logged in, you can start using PlusVasis to manage your tasks and projects.

![](https://hedgedoc.plusvasis.xyz/uploads/0b3e5fd1-b944-41d6-b030-20df32de0d28.png)

Now that you’re logged in you’ll be greeted with your dashboard:

![](https://hedgedoc.plusvasis.xyz/uploads/6d4cc493-cc6e-4631-a329-4db2225ac23c.png)

From your dashboard, you can start managing your tasks and projects using PlusVasis. You'll also see a navigation bar at the top of the screen with four options: 'Home', 'About', 'FAQ', and 'Sign Out'. Clicking on 'Home' will take you back to your dashboard. 'About' will provide more information about our application, and 'FAQ' will provide answers to frequently asked questions. If you need to sign out, simply select 'Sign Out' from the navigation bar.

Here's the about page:

![](https://hedgedoc.plusvasis.xyz/uploads/f20c7bc4-af3b-492f-9c5c-c251f2498340.png)

Here's the FAQ page:

![](https://hedgedoc.plusvasis.xyz/uploads/dd33ef21-f557-4be1-a9d6-fb668588649d.png)

In addition to the navigation bar, you'll also see a button labelled 'Create Container' on your dashboard. This button allows you to create your first container and start managing it with PlusVasis. If you haven't created any containers yet, you'll see a message on your dashboard indicating that there are no containers to display. Once you create containers, they'll be listed here on your dashboard for easy access.

With PlusVasis, managing your tasks and projects is a breeze. Our streamlined interface and intuitive navigation make it easy to get started and stay organised.

### Creating a Container

Ready to create your first container? Great! To get started, simply click the 'Create Container' button on your dashboard. This will bring up a form where you can specify the details of your container.

![](https://hedgedoc.plusvasis.xyz/uploads/1c57482d-faa3-4a13-ae23-e1264ef6e713.png)

First, give your container a name. This can be any label you choose, and it will be used to identify the container in your list of containers.

Next, you'll need to select the Docker image you wish to pull from a registry. You can choose from a variety of popular images, or specify a custom image if you prefer.

You'll also have the option to select a shell, which will be used to run your container. This can be any shell you're comfortable working with, such as Bash or Zsh.

![](https://hedgedoc.plusvasis.xyz/uploads/6317f707-dd37-4aea-b0c3-babef9ee14b2.png)

If you need to specify volumes or environment variables for your container, you can do so using the appropriate fields in the form. You can also specify the port number that your container should use.

If you want to make your container publicly accessible, you can select the 'Expose' option. This will allow anyone to access your container via its public IP address.

Finally, you can specify the CPU and Memory requirements for your container, using the range of 100-1000 and 300-2000 respectively.
The default values work well for most containers!

Once you've entered all of your container specifications, click the 'Create Container' button to create your container. You'll be redirected back to your dashboard, where you'll see your newly created container listed with its name.

![](https://hedgedoc.plusvasis.xyz/uploads/9322e36a-a017-4aac-bc12-8fbb50b867dc.png)

Congratulations! You've successfully created your first container with PlusVasis.

### Docker-Compose Import (For Advanced Users)

For more advanced users, we offer the option to create a container using a Docker Compose file. To access this feature, simply navigate to the Create Container page and click on the 'Import from docker-compose' button located at the bottom. This will take you to the Docker Compose page where you'll find an editor that allows you to write or paste in your own Docker Compose file.

![](https://hedgedoc.plusvasis.xyz/uploads/bf81b044-2a6a-4d5f-814a-8842f22aae67.png)

![](https://hedgedoc.plusvasis.xyz/uploads/678be0da-7a56-400a-bcd4-c316fcbf0987.png)

With this powerful feature, you can specify multiple containers, their images, and how they interact with each other, all in a single file. Once you've written your Docker Compose file, simply press the 'Submit' button and your container(s) will be created.

It's important to note that this feature is intended for advanced users with knowledge of Docker Compose syntax and configuration. If you're not familiar with Docker Compose, we recommend sticking to the standard container creation method described earlier.

## Container Options

### Interacting with a Container

You can interact with your container with the built-in terminal on the "Shell" tab of the Container Options page. Use the terminal to run commands, edit files and transfer data.

For example, you can use git to clone your website straight into your container!

![](https://hedgedoc.plusvasis.xyz/uploads/a72e4d25-8f9e-49e2-a2bd-fc242f10b18d.png)

### View Container Logs

As a user, you have the option to view the logs of your container, including both standard output and standard error. To access these logs, simply navigate to the container page for the container you wish to view. Once there, you'll see three tabs under the heading:

- Shell
- Logs
- Settings

To view the logs, click on the "Log" tab. By default, you'll see the standard output logs, but you can switch to viewing the standard error logs by selecting the dropdown menu and choosing "STDERR".

Here's an example of what the standard output logs look like:

![](https://hedgedoc.plusvasis.xyz/uploads/64b8715f-f8da-4eff-a543-438dcd937485.png)

And here's an example of what the standard error logs look like:

![](https://hedgedoc.plusvasis.xyz/uploads/c4e6ec42-a207-4250-b6b7-441229c81c35.png)

With this feature, you can easily monitor the activity of your container and troubleshoot any issues that may arise.

### Updating a Container

To update a container, simply navigate to the specific container page and click on the "Settings" tab. This will display the container specifications that you chose during its creation. From here, you can update any of the details based on the options provided. For example, you can change the Docker image, volumes, environment variables, port number, and CPU/memory specifications. Once you have made your desired changes, simply click on the "Update Container" button to submit them. Keep in mind that any changes made to a container will be reflected in real-time and may affect its performance, so it's always a good idea to double-check everything before clicking on the update button.

![](https://hedgedoc.plusvasis.xyz/uploads/d24019ae-0b6a-4de3-a3ab-471f2ffe4a16.png)

### Stopping a Container

Stopping a container is a straightforward process. Simply navigate to the container page and locate the button panel on the right-hand side of the page, below the navigation bar. The button on the middle right is labelled "Stop". Click on it to stop the container. Once you do so, the container will immediately stop running, and its status will change from "Running" to "Stopped". If you wish to start the container again later, you can do so by clicking the "Start" button on the same button panel.

![](https://hedgedoc.plusvasis.xyz/uploads/518856d2-de49-410f-be04-eda9943fd841.png)


### Starting a Container

To start a container, navigate to the container page and you'll see a list of four buttons on the right-hand side of the page underneath the navigation bar. Click the start button to start the container that is currently stopped.

![](https://hedgedoc.plusvasis.xyz/uploads/e5bb1933-bf69-454d-8426-63df5b25d097.png)

### Restarting a Container

To restart a container, navigate to the container page and you'll see a list of four buttons on the right-hand side of the page underneath the navigation bar. Click the restart button to stop and start the container again.

![](https://hedgedoc.plusvasis.xyz/uploads/fed4b3ef-b482-47c1-bd5c-dfd63b0ec09f.png)

### Deleting a Container

To delete a container, navigate to the container page and you'll see a list of four buttons on the right-hand side of the page underneath the navigation bar. Click the delete button to remove the container entirely. Be careful with this option as once you delete the container, you will lose all the data inside it.

![](https://hedgedoc.plusvasis.xyz/uploads/ae05ed2f-a2b1-464d-bbf7-bfee93620759.png)

## Container Configuration

### Exposing a container

To expose your container, specify the port you'd like to expose and enable the expose option. The container will then be accessible at `https://<your-username>-<your-container-name>.plusvasis.xyz` and a link to this can be found beside the container name on the container dashboard.

### Inter-container communication

You can reference other containers with the `{{containerName}}` syntax in your environment variables. For example, you can have an environment variable containing the connection address to a Postgres database container like: `DB_URL=psql://{{postgres}}`
