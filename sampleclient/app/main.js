import axios from 'axios' 
import grpcWeb from 'grpc-web';
import {WorldSyncRenderClient} from './protos/world_render_grpc_web_pb';
import {WorldRenderRequest, World} from './protos/world_render_pb';

let app = (name) => {
    	console.log(`hello from ${name}`);
    	var world;
	//axios.get('http://localhost:8000/world')
        //	.then(function (response) {
        //  		response.data.map(function (world) {
	//			world = world
	//		})
	//	});

	function drawWorld(world) {
		var canvas = document.getElementById("main");
    		var ctx = canvas.getContext("2d");
		for (var i in world.getEntityList()) {
			var entity = world.getEntityList()[i]
    			if (entity.getType() === "villager") {
				ctx.fillStyle = "#FF0000";
			} else if (entity.getType() === "grass") {
				ctx.fillStyle = "#00FF00";
			}
			var area = entity.getArea();
			ctx.fillRect(area.getX(), area.getY(), area.getWidth(), area.getHeight());
		}
	}

	var worldSync = new WorldSyncRenderClient(
  			'http://localhost:8080');
	var request = new WorldRenderRequest();
	var stream = worldSync.worldStartRender(request);

	stream.on('data', function(response) {
  		console.log("world received");
		drawWorld(response);
	});
	stream.on('status', function(status) {
  		console.log(status.code);
  		console.log(status.details);
  		console.log(status.metadata);
	});
	stream.on('end', function(end) {
  		// stream end signal
	});
    	
}
app('distgame');
